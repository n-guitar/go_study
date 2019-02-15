package main

import "fmt"

func main() {

	var a = 1
	// 足し算
	a += 1
	fmt.Printf("足し算の問題です。1 + 1 = %v\n", a)
	v := a

	//　引き算
	a -= 1
	fmt.Printf("引き算の問題です。%v - 1 = %v\n", v, a)
	v = a

	//　掛け算
	a *= 2
	fmt.Printf("掛け算の問題です。%v × 2 = %v\n", v, a)
	v = a

	//　割り算
	a /= 2
	fmt.Printf("割り算の問題です。%v ÷ 2 = %v\n", v, a)

}
