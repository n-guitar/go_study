// liner search
package main

import (
	"fmt"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	key := 1

	found := false
	for _, v := range numbers {
		if v == key {
			found = true
		}
	}

	if found {
		fmt.Println("found!")
	} else {
		fmt.Println("not found...")
	}
}
