package main

import (
	"fmt"
	"math"
)

func main() {

	// number must be begger than -1.
	numbers := []int{2, 3, 4, 5, 100, 2, 55, 99, 0}

	min := math.MaxInt64
	max := -1

	for _, v := range numbers {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Printf("min : %d\n", min)
	fmt.Printf("max : %d\n", max)
}
