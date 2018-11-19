package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")

	result, err := cmd.Output()

	if err != nil {
		fmt.Printf("エラーが発生:%s", err)
		os.Exit(1)
	}

	fmt.Printf("%s", result)
}
