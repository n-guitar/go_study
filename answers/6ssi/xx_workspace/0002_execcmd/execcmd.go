package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Oops -> %s", err)
		os.Exit(1)
	}
	fmt.Printf("Command returuned :\n%s", out)
}
