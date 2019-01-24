package main

import (
	"fmt"
	"strconv"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println("sum : " + strconv.Itoa(sum))
}
