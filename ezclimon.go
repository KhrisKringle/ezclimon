package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	os.Chdir("/")

	if dir, _ := os.Getwd(); dir != "/" {
		log.Fatal(fmt.Errorf("Not in root directory: %s", dir))
	}

	fmt.Printf("Welcome \033[1;36m%s\033[0m, What would you like to do?\n", os.Getenv("USER"))

	var choice string

	for choice != "5" && choice != "exit" && choice != "quit" && choice != "q" && choice != "x" {

		Boxed_Print("1. Check Storage\n2. Network Info\n3. System Info\n4. Memory Info\n5. Exit\n")
		fmt.Print("What ll it be?: ")
		//choice = "mem"
		fmt.Scanln(&choice)

		switch choice {
		case "1", "storage", "stor":
			err := Storage_Check()
			if err != nil {
				log.Fatal(err)
			}

			var stor_choice string
			fmt.Print("What else do you want to check for storage: ")
			fmt.Scanln(&stor_choice)

			switch stor_choice {
			case "check":
				fmt.Println("Storage integrity check is not implemented yet.")
				// err := Storage_Integrity_Check()
				// if err != nil {
				// 	log.Fatal(err)
				// }
			}
		case "2", "network", "net":
			err := Network_Info()
			if err != nil {
				log.Fatal(err)
			}
		case "3", "system", "sys":
			err := System_Info()
			if err != nil {
				log.Fatal(err)
			}
		case "4", "memory", "mem":
			err := Memory_Info()
			if err != nil {
				log.Fatal(err)
			}

			var mem_choice string
			fmt.Print("What else do you want to check with memory: ")
			//mem_choice = "check"
			fmt.Scanln(&mem_choice)

			switch mem_choice {
			case "check":
				err := Memory_Integrity_Check()
				if err != nil {
					log.Fatal(err)
				}
			}
		case "5", "exit", "quit", "q", "x":
			fmt.Println("Goodbye!")
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func Storage_Check() error {
	fmt.Println("\033[1;36m===Checking storage...===\033[0m")
	cmd, err := exec.Command("df", "-h").Output()
	if err != nil {
		log.Fatal(err)
	}
	Boxed_Print(string(cmd))
	fmt.Println("\033[1;36m===Checking storage Complete===\033[0m")
	return nil
}

func Network_Info() error {
	fmt.Println("\033[1;36m===Checking network info...===\033[0m")
	cmd, err := exec.Command("ip", "a").Output()
	if err != nil {
		log.Fatal(err)
	}
	Boxed_Print(string(cmd))
	fmt.Println("\033[1;36m===Checking Network Info Complete===\033[0m")
	return nil
}

func System_Info() error {
	fmt.Println("\033[1;36m===Checking system info...===\033[0m")
	cmd, err := exec.Command("uname", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}
	Boxed_Print(string(cmd))
	fmt.Println("\033[1;36m===Checking System Info Complete===\033[0m")
	return nil
}

func Memory_Info() error {
	fmt.Println("\033[1;36m===Checking memory info...===\033[0m")
	cmd, err := exec.Command("free", "-h").Output()
	if err != nil {
		log.Fatal(err)
	}
	Boxed_Print(string(cmd))
	fmt.Println("\033[1;36m===Checking Memory Info Complete===\033[0m")
	return nil
}

func Memory_Integrity_Check() error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	stopChan := make(chan struct{})
	go func() {
		// This will block until a newline is entered.
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		close(stopChan)
	}()

	for {
		// Clear screen
		fmt.Print("\033[H\033[2J")
		fmt.Println("\033[1;36m===Running Memory Integrity Audit... (Press Enter to stop)===\033[0m")

		file, err := os.Open("/proc/meminfo")
		if err != nil {
			return err
		}

		var memTotal, memAvailable float64

		// Parse /proc/meminfo line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Fields(line)

			if len(fields) < 2 {
				continue
			}

			switch fields[0] {
			case "MemTotal:":
				val, err := strconv.ParseFloat(fields[1], 64)
				if err != nil {
					file.Close()
					return fmt.Errorf("failed to parse MemTotal: %v", err)
				}
				memTotal = val
			case "MemAvailable:":
				// Available is better than Free, as it accounts for reclaimable cache
				val, err := strconv.ParseFloat(fields[1], 64)
				if err != nil {
					file.Close()
					return fmt.Errorf("failed to parse MemAvailable: %v", err)
				}
				memAvailable = val
			}
		}
		file.Close()

		if err := scanner.Err(); err != nil {
			return err
		}

		// Calculate memory utilization percentage
		usedPercent := ((memTotal - memAvailable) / memTotal) * 100

		var resultMessage string
		if usedPercent >= 90.0 {
			resultMessage = fmt.Sprintf("🚨 \033[1;31mCRITICAL:\033[1;0m Memory usage at \033[1;31m%.2f%%!\033[1;0m\nSystem is at risk of OOM kills.\nPress Enter to stop.", usedPercent)
		} else if usedPercent >= 75.0 {
			resultMessage = fmt.Sprintf("⚠️ \033[1;33mWARNING:\033[1;0m Memory usage high at \033[1;33m%.2f%%\033[1;0m.\nPress Enter to stop.", usedPercent)
		} else {
			resultMessage = fmt.Sprintf("✅ \033[1;32mHEALTHY:\033[1;0m Memory usage normal at \033[1;32m%.2f%%.\033[1;0m\nPress Enter to stop.", usedPercent)
		}

		Boxed_Print(resultMessage)

		select {
		case <-stopChan:
			fmt.Print("\033[H\033[2J") // Clear screen
			return nil
		case <-ticker.C:
			// continue loop
		}
	}
}

func Boxed_Print(text string) {
	lines := strings.Split(strings.TrimSpace(text), "\n")

	// Regex to strip ANSI escape codes for correct width calculation
	re := regexp.MustCompile("\x1b\\[[0-9;]*m")

	// Find max line width
	maxWidth := 0
	for _, line := range lines {
		plainLine := re.ReplaceAllString(line, "")
		if len(plainLine) > maxWidth {
			maxWidth = len(plainLine)
		}
	}

	// Draw box
	fmt.Println("\n╔" + strings.Repeat("═", maxWidth+2) + "╗")
	for _, line := range lines {
		plainLine := re.ReplaceAllString(line, "")
		padding := strings.Repeat(" ", maxWidth-len(plainLine))
		fmt.Printf("║ %s%s ║\n", line, padding)
	}
	fmt.Println("╚" + strings.Repeat("═", maxWidth+2) + "╝")
}
