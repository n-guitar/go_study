// object sort
package main

import (
	"fmt"
)

type User struct {
	name  string
	score int
}

func main() {

	var xs = []User{
		{"a", 28},
		{"b", 25},
		{"c", 29},
		{"d", 11},
		{"e", 12},
		{"f", 13},
		{"g", 22},
		{"h", 17},
		{"i", 1}}

	for i := 0; i < len(xs); i++ {
		for j := 0; j < len(xs)-1-i; j++ {
			if xs[j].score > xs[j+1].score {
				xs[j], xs[j+1] = xs[j+1], xs[j]
			}
		}
	}

	fmt.Println(xs)

}
