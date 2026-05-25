package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	// os.Chdir("/")

	// if dir, _ := os.Getwd(); dir != "/" {
	// 	log.Fatal(fmt.Errorf("Not in root directory: %s", dir))
	// }

	fmt.Printf("Welcome \033[1;36m%s\033[0m, What would you like to do?\n", os.Getenv("USER"))

	for {
		var choice string
		Boxed_Print("1. 💾 Check Storage\n2. 🌐 Network Info\n3. 💻 System Info\n4. 🧠 Memory Info\n5. 🚪 Exit\n")
		fmt.Print("What ll it be?: ")
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
					log.Println(err) // Use Println to avoid exiting on integrity check error
				}
			}
		case "5", "exit", "quit", "q", "x":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
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
