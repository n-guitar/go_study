package main

import "fmt"

// ToDo 引数によって型を変更する。
// 例えば、1+1=2 1.0+1.0= 2.0 と表示したい
func Arithmetic(a int, b int) (int, int, int, float64) {
	return a + b, a - b, a * b, float64(a) / float64(b)
}

func main() {
	add, aub, mul, div := Arithmetic(1, 1)
	fmt.Printf("足し算：%d 引き算：%d 掛け算：%d 割り算：%e \n", add, aub, mul, div)
	fmt.Printf("%d:%d:%d:%e \n", add, aub, mul, div)
}
