package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"strconv"
	"bufio
)

func main() {

	os.Chdir("/")

	if dir, _ := os.Getwd(); dir != "/" {
		log.Fatal(fmt.Errorf("Not in root directory: %s", dir))
	}

	fmt.Printf("Welcome \033[1;36m%s\033[0m, What would you like to do?\n", os.Getenv("USER"))
	
	var choice int

	for choice != 5 {
	
		Boxed_Print("1. Check Storage\n2. Network Info\n3. System Info\n4. Memory Info\n5. Exit\n")
		fmt.Print("What'll it be?: ")

	
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			err := Storage_Check()
			if err != nil {
				log.Fatalf(err)
			}
			var stor_choice string
			fmt.Print("What else do you want to check for storage: ")
			fmt.Scanln(&stor_choice)
			if stor_choice == "Inode check" {
				//Inode_Check()	
			}
		case 2:
			err := Network_Info()
			if err != nil {
				log.Fatal(err)
			}
		case 3:
			err := System_Info()
			if err != nil {
				log.Fatal(err)
			}
		case 4:
			err := Memory_Info()
			if err != nil {
				log.Fatal(err)
			}
			var mem_choice string
			fmt.Print("What else do you want to check for storage: ")
			fmt.Scanln(&mem_choice)
			if mem_choice == "integ check" {
				err := Memory_Integrity_Check()
				if err != nil {
					log.Fatal(err)
				}
			}
		case 5:
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
		return log.Fatal(err)
	}
	Boxed_Print(string(cmd))
	return fmt.Println("\033[1;36m===Checking Network Info Complete===\033[0m")
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
		return log.Fatal(err)
	}
	Boxed_Print(string(cmd))
	fmt.Println("\033[1;36m===Checking Memory Info Complete===\033[0m")
	return nil
}

func Memory_Integrity_Check() error {
	fmt.Println("\033[1;36m===Running Memory Integrity Audit...===\033[0m")

	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return log.Fatal(err)
	}
	defer file.Close()

	var memTotal, memAvailable float64

	// Parse /proc/meminfo line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		
		if len(fields) < 2 {
			continue
		}

		if fields[0] == "MemTotal:" {
			val, _ := strconv.ParseFloat(fields[1], 64)
			memTotal = val
		} else if fields[0] == "MemAvailable:" {
			// Available is better than Free, as it accounts for reclaimable cache
			val, _ := strconv.ParseFloat(fields[1], 64)
			memAvailable = val
		}
	}

	// Calculate memory utilization percentage
	usedPercent := ((memTotal - memAvailable) / memTotal) * 100

	var resultMessage string
	if usedPercent >= 90.0 {
		resultMessage = fmt.Sprintf("🚨 \033[1;31mCRITICAL:\033[1;31m Memory usage at \033[1;31m%.2f%%!\033[1;31m\nSystem is at risk of OOM kills.", usedPercent)
	} else if usedPercent >= 75.0 {
		resultMessage = fmt.Sprintf("⚠️ \033[1;33mWARNING:\033[1;33m Memory usage high at \033[0;33m%.2f%%\033[0;33m.", usedPercent)
	} else {
		resultMessage = fmt.Sprintf("✅ \033[1;32mHEALTHY:\033[1;32m Memory usage normal at %.2f%%.", usedPercent)
	}

	Boxed_Print(resultMessage)
	return nil
}

func Boxed_Print(text string) {
	lines := strings.Split(strings.TrimSpace(text), "\n")

	// Find max line width
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Draw box
	fmt.Println("╔" + strings.Repeat("═", maxWidth+2) + "╗")
	for _, line := range lines {
		fmt.Printf("║ %-*s ║\n", maxWidth, line)
	}
	fmt.Println("╚" + strings.Repeat("═", maxWidth+2) + "╝")
}
