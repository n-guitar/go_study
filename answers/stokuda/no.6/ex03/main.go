package main

import (
	"fmt"
	"os"
	"strconv"
)

func plus(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}

func main() {
	var args []int
	for _, x := range os.Args[1:] {
		x_i, _ := strconv.Atoi(x)
		args = append(args, x_i)
	}
	fmt.Println(plus(args))
}
