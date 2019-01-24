package main

import (
	"fmt"
)

func main() {

	numbers := []float64{1, 2, 3, 4, 5, 8}

	var sum float64 = 0
	for _, v := range numbers {
		sum += v
	}
	avg := sum / float64(len(numbers))

	fmt.Printf("sum: %.2f, length of numbers: %d\n", sum, len(numbers))
	fmt.Printf("avg : %.2f\n", avg)
}
