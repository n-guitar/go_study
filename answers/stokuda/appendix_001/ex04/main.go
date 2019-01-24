package main

import (
	"fmt"
)

func main() {

	x := 10
	y := 20
	fmt.Printf("x: %d, y: %d\n", x, y)

	x, y = y, x
	fmt.Printf("x: %d, y: %d\n", x, y)

	/*
		x := 10
		y := 20
		tmp := 30

		tmp = x
		x = y
		y = tmp

	*/

}
