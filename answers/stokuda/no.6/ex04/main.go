package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func _sha256(in string) [32]byte {
	return sha256.Sum256([]byte(in))
}

func _sha384(in string) [48]byte {
	return sha512.Sum384([]byte(in))
}

func _sha512(in string) [64]byte {
	return sha512.Sum512([]byte(in))
}

func get_in() string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	in := stdin.Text()
	if err := stdin.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return in
}

func main() {
	var sha string
	flag.StringVar(&sha, "s", "", "-s (sha256|sha384|sha512)　を引数に指定")
	flag.Parse()
	sha_flag := os.Args[2]

	switch sha_flag {
	case "sha256":
		fmt.Printf("sha256: %x", _sha256(get_in()))
	case "sha384":
		fmt.Printf("sha384: %x", _sha384(get_in()))
	case "sha512":
		fmt.Printf("sha512: %x", _sha512(get_in()))
	default:
		fmt.Println("-s (sha256|sha384|sha512)　を引数に指定")
	}

}
