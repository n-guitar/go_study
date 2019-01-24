package main

import (
	"fmt"
)

// bubble sort
func main() {

	xs := []int{28, 25, 29, 11, 12, 13, 22, 17, 1}

	for i := 0; i < len(xs); i++ {
		for j := 0; j < len(xs)-1-i; j++ {
			if xs[j] > xs[j+1] {
				xs[j], xs[j+1] = xs[j+1], xs[j]
			}
		}
	}

	fmt.Println(xs)

}
