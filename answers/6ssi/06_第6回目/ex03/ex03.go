package main

import (
	"fmt"
	"os"
	"strconv"
)

func arithmetic(intargs []int) (int, int, int, float64) {
	plus := intargs[0]
	minus := intargs[0]
	multiple := intargs[0]
	divide := float64(intargs[0])
	for i := 1; i < len(intargs); i++ {
		plus = plus + intargs[i]
	}
	for i := 1; i < len(intargs); i++ {
		minus = minus - intargs[i]
	}
	for i := 1; i < len(intargs); i++ {
		multiple = multiple * intargs[i]
	}
	for i := 1; i < len(intargs); i++ {
		divide = divide / float64(intargs[i])
	}
	return plus, minus, multiple, divide
}

func convarg(inputargs []string) []int {
	var outargs []int
	for i := 1; i < len(inputargs); i++ {
		inputstr, err := strconv.Atoi(inputargs[i])
		if err != nil {
			fmt.Printf("Failed to convert %s to Int\n", inputargs[i])
		}
		outargs = append(outargs, inputstr)
	}
	return outargs
}

func main() {
	plus, minus, multiple, divide := arithmetic(convarg(os.Args))
	fmt.Printf("+ : %d\n", plus)
	fmt.Printf("- : %d\n", minus)
	fmt.Printf("× : %d\n", multiple)
	fmt.Printf("÷ : %g\n", divide)
}
