package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(`./all.txt`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open err!: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "file err!: %v\n", sc.Err()) //こういう書き方で良いかわからない
			os.Exit(1)
		}
		fmt.Printf("%d行目: %s\n", i, sc.Text())
	}

}
