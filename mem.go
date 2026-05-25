package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

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

			switch fields[0] {
			case "MemTotal:":
				val, err := strconv.ParseFloat(fields[1], 64)
				if err != nil {
					return fmt.Errorf("failed to parse MemTotal: %w", err)
				}
				memTotal = val
			case "MemAvailable:":
				// Available is better than Free, as it accounts for reclaimable cache
				val, err := strconv.ParseFloat(fields[1], 64)
				if err != nil {
					return fmt.Errorf("failed to parse MemAvailable: %w", err)
				}
				memAvailable = val
			}
		}

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
