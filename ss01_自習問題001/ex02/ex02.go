package main

import "fmt"

func main() {
	slice := []int{2, 4, 10, 1}
	sum := 0
	for _, g := range slice {
		sum += g
	}
	l := len(slice)
	sum = sum / l
	fmt.Printf("sum:%d\n", sum)
}
