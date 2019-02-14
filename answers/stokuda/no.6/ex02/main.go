package main

import "fmt"

func four_op(a int, b int) {
	fmt.Println(add(a, b))
	fmt.Println(minus(a, b))
	fmt.Println(multiply(a, b))
	fmt.Println(divide(a, b))
}

func add(x int, y int) int {
	return x + y
}
func minus(x int, y int) int {
	return x - y
}
func multiply(x int, y int) int {
	return x * y
}
func divide(x int, y int) int {
	return x / y
}

func main() {
	four_op(4, 2)
}
