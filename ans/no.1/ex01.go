package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-a", "-l")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Command failed.")
		panic("exit")
	}
	fmt.Println(string(out))
}
