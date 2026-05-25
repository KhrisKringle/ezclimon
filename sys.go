package main

import (
	"fmt"
	"log"
	"os/exec"
)

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
