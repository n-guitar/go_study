package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(`./kon.txt`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open err!: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	if true != sc.Scan() {
		fmt.Fprintf(os.Stderr, "file err!: %v\n", sc.Err()) //こういう書き方で良いかわからない
		os.Exit(1)
	}
	fmt.Printf("ファイルの中身: %s\n", sc.Text())

}
