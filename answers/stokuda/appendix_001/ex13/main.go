// shuffle
package main

import (
	"fmt"
	"math/rand"
)

// bubble sort
func main() {

	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for i := 0; i < 100; i++ {
		x := rand.Intn(len(xs) - 1)
		y := rand.Intn(len(xs) - 1)
		xs[x], xs[y] = xs[y], xs[x]
	}

	fmt.Println(xs)

}
