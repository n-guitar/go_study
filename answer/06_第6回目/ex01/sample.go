package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main() {
	answer := plus(1, 2)
	fmt.Println(answer)

}
