package main

import (
	"fmt"
)

func main() {
	for index := 10; index > 0; index-- {
		if index == 5 {
			break
		}
		fmt.Println(index)
	}
}
