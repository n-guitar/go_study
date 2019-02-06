package main

import "fmt"

func arithmetic(a int, b int) (int, int, int, int) {
	plus := a + b
	minus := a - b
	multiple := a * b
	divide := a / b
	return plus, minus, multiple, divide
}

func main() {
	plus, minus, multiple, divide := arithmetic(10, 2)
	fmt.Printf("+ : %d\n", plus)
	fmt.Printf("- : %d\n", minus)
	fmt.Printf("× : %d\n", multiple)
	fmt.Printf("÷ : %d\n", divide)
}
