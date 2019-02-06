package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var sha string
	stdin := bufio.NewScanner(os.Stdin)
	flag.StringVar(&sha, "s", "", "-s (sha256|sha384|sha512)　を引数に指定")
	flag.Parse()
	sha_flag := os.Args[2]
	switch sha_flag {
	case "sha256":
		for stdin.Scan() {
			if err := stdin.Err(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			sha_out := sha256.Sum256([]byte(stdin.Text()))
			fmt.Printf("%x\n", sha_out)
		}
		fmt.Print("sha256", stdin, "\n")
	case "sha384":
		for stdin.Scan() {
			if err := stdin.Err(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			sha_out := sha512.Sum384([]byte(stdin.Text()))
			fmt.Printf("%x\n", sha_out)
		}
		fmt.Print("sha256", stdin, "\n")
	case "sha512":
		for stdin.Scan() {
			if err := stdin.Err(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			sha_out := sha512.Sum512([]byte(stdin.Text()))
			fmt.Printf("%x\n", sha_out)
		}
		fmt.Print("sha512", stdin, "\n")
	}
	fmt.Println("-s (sha256|sha384|sha512)　を引数に指定")
}
