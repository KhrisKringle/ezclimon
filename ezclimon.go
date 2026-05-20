package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	os.Chdir("/")

	if dir, _ := os.Getwd(); dir != "/" {
		log.Fatal(fmt.Errorf("Not in maint directory: %s", dir))
	}

	fmt.Printf("Welcome %s, What would you like to do?\n", os.Getenv("USER"))
	fmt.Println("1. Check Storage")
	fmt.Println("2. Network Info")
	fmt.Println("3. System Info")
	fmt.Println("4. Exit")
	fmt.Print("What'll it be?: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		Storage_Check()
	case 2:
		Network_Info()
	case 3:
		System_Info()
	case 4:
		fmt.Println("Goodbye!")
	default:
		fmt.Println("Invalid choice, please try again.")
	}
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

func Storage_Check() {
	fmt.Println("\033[1;36m===Checking storage...===\033[0m")
	cmd, err := exec.Command("df", "-h").Output()
	if err != nil {
		log.Fatal("Could not run the Command for some reason: ", err)
	}
	Boxed_Print(string(cmd))
}

func Network_Info() {
	fmt.Println("\033[1;36m===Checking network info...===\033[0m")
	cmd, err := exec.Command("ip", "a").Output()
	if err != nil {
		log.Fatal("Could not run the Command for some reason: ", err)
	}
	Boxed_Print(string(cmd))
}

func System_Info() {
	fmt.Println("\033[1;36m===Checking system info...===\033[0m")
	cmd, err := exec.Command("uname", "-a").Output()
	if err != nil {
		log.Fatal("Could not run the Command for some reason: ", err)
	}
	Boxed_Print(string(cmd))
}
