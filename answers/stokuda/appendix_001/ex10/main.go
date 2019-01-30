package main

import (
	"fmt"
)

// shell sort
func main() {

	xs := []int{19, 11, 10, 7, 8, 9, 17, 18, 20, 4, 3, 15, 16, 1, 5, 14, 6, 2, 13, 12}

	for h := 4; h > 0; {
		for i := 1; i < len(xs); i++ {
			for j := i; j >= h; j -= h {
				if xs[j-h] < xs[j] {
					break
				}
				xs[j-h], xs[j] = xs[j], xs[j-h]
			}
		}
		h /= 2
	}

	fmt.Println(xs)
}
