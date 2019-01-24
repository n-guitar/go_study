// binary search
package main

import (
	"fmt"
)

func bin_search(numbers []int, key int, left int, right int) int {
	if right < left {
		return -1
	}

	mid := (left + right) / 2
	if numbers[mid] == key {
		return mid
	} else if numbers[mid] > key {
		return bin_search(numbers, key, left, mid-1)
	} else if numbers[mid] < key {
		return bin_search(numbers, key, mid+1, right)
	}
	return -999
}

func main() {

	numbers := []int{1, 3, 5, 11, 12, 13, 17, 22, 25, 28}
	key := 11

	found := bin_search(numbers, key, 0, len(numbers)-1)
	if found != -1 {
		fmt.Println("found!")
		fmt.Println(found + 1)
	} else {
		fmt.Println("not found...")
	}

}
