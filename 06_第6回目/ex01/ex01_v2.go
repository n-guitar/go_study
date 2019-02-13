package main

import "fmt"

// 参照：https://qiita.com/high5/items/4e2580241039c950e1c4
func calc(a int, b int) (int, int, int, int) {
	plus := a + b
	minus := a - b
	multi := a * b
	divi := a / b
	return plus, minus, multi, divi
}

func main() {
	a := 1
	b := 2
	plus, minus, multi, divi := calc(a, b)
	fmt.Println(plus)
	fmt.Println(minus)
	fmt.Println(multi)
	fmt.Println(divi)
}
