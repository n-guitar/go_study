package main

import (
	"fmt"
)

// insert sort
func main() {

	xs := []int{19, 11, 10, 7, 8, 9, 17, 18, 20, 4, 3, 15, 16, 1, 5, 14, 6, 2, 13, 12}

	for i := 1; i < len(xs); i++ {
		for j := i; j > 0; j-- {
			if xs[j-1] < xs[j] {
				break
			}
			xs[j-1], xs[j] = xs[j], xs[j-1]
		}
	}

	fmt.Println(xs)
}
