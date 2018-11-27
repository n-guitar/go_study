package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.Open("./w.txt")
	defer f.Close()
	if err != nil {
		fmt.Println("file not found")
		os.Exit(1)
	}

	r := bufio.NewReaderSize(f, 1024)
	for line, err := r.ReadString('\n'); err == nil; line, err = r.ReadString('\n') {
		fmt.Println(line)
		time.Sleep(1 * time.Second)
	}

}
