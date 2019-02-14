package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func s256(intxt string) [32]byte {
	var sha32 [32]byte
	sha32 = sha256.Sum256([]byte(intxt))
	return sha32
}

func s384(intxt string) [48]byte {
	var sha48 [48]byte
	sha48 = sha512.Sum384([]byte(intxt))
	return sha48
}

func s512(intxt string) [64]byte {
	var sha64 [64]byte
	sha64 = sha512.Sum512([]byte(intxt))
	return sha64
}

func main() {
	var sha string
	flag.StringVar(&sha, "s", "sha256", "-s (sha256|sha384|sha512)　を引数に指定")
	flag.Parse()
	shaflag := os.Args[2]
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		if err := scn.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		switch shaflag {
		case "sha256":
			fmt.Printf("sha256 :%x\n", s256(scn.Text()))
		case "sha384":
			fmt.Printf("sha384 :%x\n", s384(scn.Text()))
		case "sha512":
			fmt.Printf("sha512 :%x\n", s512(scn.Text()))
		}
	}
}
