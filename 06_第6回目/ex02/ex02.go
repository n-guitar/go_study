package main

import "fmt"

func arithmetic(a int, b int) (int, int, int, int) {
	plus := a + b
	minus := a - b
	multiple := a * b
	divide := a / b
	return plus, minus, multiple, divide
}

func pluschk(a int, b int, target int) string {
	plus := a + b
	output := "NG"
	if plus == target {
		output = "OK"
	}
	return output
}

func minuschk(a int, b int, target int) string {
	minus := a - b
	output := "NG"
	if minus == target {
		output = "OK"
	}
	return output
}

func multiplechk(a int, b int, target int) string {
	multiple := a * b
	output := "NG"
	if multiple == target {
		output = "OK"
	}
	return output
}

func dividechk(a int, b int, target int) string {
	divide := a / b
	output := "NG"
	if divide == target {
		output = "OK"
	}
	return output
}

func main() {
	plus, minus, multiple, divide := arithmetic(10, 2)
	fmt.Printf("+ : %d  / Check : %s\n", plus, pluschk(10, 2, plus))
	fmt.Printf("- : %d  / Check : %s\n", minus, minuschk(10, 2, minus))
	fmt.Printf("× : %d  / Check : %s\n", multiple, multiplechk(10, 2, multiple))
	fmt.Printf("÷ : %d  / Check : %s\n", divide, dividechk(10, 2, divide))
}
