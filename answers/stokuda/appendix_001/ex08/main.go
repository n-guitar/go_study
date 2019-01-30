package main

import (
	"fmt"
)

// selection sort.
func main() {

	xs := []int{28, 25, 29, 11, 12, 13, 22, 17, 1}

	for i := 0; i < len(xs); i++ {
		min := -1
		for j := i; j < len(xs); j++ {
			if min == -1 || xs[min] > xs[j] {
				min = j
			}
		}
		xs[i], xs[min] = xs[min], xs[i]
	}

	fmt.Println(xs)

}
