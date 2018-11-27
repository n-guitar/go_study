package main

import (
	"fmt"
)

func main() {
	s := "Hello"
	c := []byte(s) // 文字列sを[]byte型にキャストする
	c[0] = 'c'
	s2 := string(c) // cをstring型にキャスト
	fmt.Printf("%s\n", s2)

	s3 := "H" + s[1:]
	fmt.Printf("%s\n", s3)

}
