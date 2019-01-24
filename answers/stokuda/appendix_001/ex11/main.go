// quick sort.
package main

import (
	"fmt"
)

func main() {

	//	xs := []int{19, 11, 10, 7, 8, 9, 17, 18, 20, 4, 3, 15, 16, 1, 5, 14, 6, 2, 13, 12}
	//	xs := []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 22, 24, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	xs := []int{8, 4, 3, 7, 6, 5, 2, 1}
	quick_sort(xs, 0, len(xs)-1)
	fmt.Println(xs)
}

func med3(x int, y int, z int) int {
	if x < y {
		if y < z {
			return y
		} else if z < x {
			return x
		} else {
			return z
		}
	} else {
		if z < y {
			return y
		} else if x < z {
			return x
		} else {
			return z
		}
	}
}

func quick_sort(xs []int, l int, r int) {

	if l < r {
		i := l
		j := r
		p := med3(xs[i], xs[i+(j-i)/2], xs[j])
		for {

			for ; xs[i] < p; i++ {
			}

			for ; xs[j] > p; j-- {
			}

			if i >= j {
				break
			}

			xs[i], xs[j] = xs[j], xs[i]
			i++
			j--
		}
		fmt.Println(xs)
		fmt.Println(l, r, p)
		quick_sort(xs, l, i-1)
		quick_sort(xs, j+1, r)

	}

}
