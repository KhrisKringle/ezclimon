package main

import (
	"fmt"
	"log"
	"os/exec"
)

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
