package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// コマンドの引数は、,区切りで指定
	cmd := exec.Command("docker", "images")

	result, err := cmd.Output()

	if err != nil {
		fmt.Printf("エラーが発生:%s", err)
		os.Exit(1)
	}

	fmt.Printf("%s", result)
}
