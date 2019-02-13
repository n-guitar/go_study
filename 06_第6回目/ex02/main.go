package main

import "fmt"

// ToDo 引数によって型を変更する。
// 例えば、1+1=2 1.0+1.0= 2.0 と表示したい
func Arithmetic(a float64, b float64) (float64, float64, float64, float64) {
	return a + b, a - b, a * b, a / b
}

func main() {
	add, aub, mul, div := Arithmetic(4, 2)
	fmt.Printf("足し算：%e 引き算：%e 掛け算：%e 割り算：%e \n", add, aub, mul, div)
}
