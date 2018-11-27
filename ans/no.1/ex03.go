package main

import (
	"flag"
	"fmt"
)

//
// $ go run ex03.go -name Tom -age 20
//
func main() {
	var (
		n = flag.String("name", "anonymouse", "name flag")
		a = flag.Int("age", 0, "age flag")
	)
	flag.Parse()
	fmt.Println(*n, *a)
}
