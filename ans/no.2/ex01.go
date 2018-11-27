package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./w.txt")
	defer f.Close()
	if err != nil {
		fmt.Println("file not found")
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	n, _ := f.Read(buf)

	fmt.Print(string(buf[:n]))
}
