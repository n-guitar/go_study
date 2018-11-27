package main

import (
	"fmt"
	"os"
)

func main() {
	for _, v := range os.Args {
		fmt.Println(v)
	}
}
