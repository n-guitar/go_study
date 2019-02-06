package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func s256() {
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		if err := scn.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		shaout := sha256.Sum256([]byte(scn.Text()))
		fmt.Printf("%x\n", shaout)
	}
}

func s384() {
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		if err := scn.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		shaout := sha512.Sum384([]byte(scn.Text()))
		fmt.Printf("%x\n", shaout)
	}
}

func s512() {
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		if err := scn.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		shaout := sha512.Sum512([]byte(scn.Text()))
		fmt.Printf("%x\n", shaout)
	}
}

func main() {
	var sha string
	stdin := bufio.NewScanner(os.Stdin)
	flag.StringVar(&sha, "s", "", "-s (sha256|sha384|sha512)　を引数に指定")
	flag.Parse()
	shaflag := os.Args[2]
	switch shaflag {
	case "sha256":
		s256()
		fmt.Print("sha256", stdin, "\n")
	case "sha384":
		s384()
		fmt.Print("sha256", stdin, "\n")
	case "sha512":
		s512()
		fmt.Print("sha512", stdin, "\n")
	}
	fmt.Println("-s (sha256|sha384|sha512)　を引数に指定")
}
