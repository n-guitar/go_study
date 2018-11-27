package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	flag1 := flag.String("s", "ps", "-s で文字列を入力")
	flag.Parse()

	fmt.Printf("フラグ：\n%s\n", *flag1)

	cmd := exec.Command("docker", *flag1)
	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("出力：\n%s", result)
}
