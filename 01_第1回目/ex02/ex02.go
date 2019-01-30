package main

/* ex02 */

import (
	"fmt"
	"os"
)

func main() {
	countargs(os.Args)
	printargs(os.Args)
}

func countargs(inputargs []string) {
	fmt.Printf("You have "+"%d"+" arguments.\n", len(inputargs[1:]))
	/* inputargs[1:]
	can also be
		inputargs[1 : len(inputargs)] */
}

func printargs(inputargs []string) {
	fmt.Printf("Here are your arguments:\n")
	for i := 1; i < len(inputargs); i++ {
		fmt.Printf("%s\n", inputargs[i])
	}
}
