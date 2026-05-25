package main

import (
	"fmt"
	"log"
	"os/exec"
)

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
