package main

/* ex03 */

import (
	"flag"
	"fmt"
)

func main() {
	inputname := flag.String("name", "gopher", "Type your name")
	inputage := flag.Int("age", 11, "Type your Age")
	flag.Parse()
	fmt.Printf("Your name is %s\n", *inputname)
	fmt.Printf("Your age is %d\n", *inputage)
}
