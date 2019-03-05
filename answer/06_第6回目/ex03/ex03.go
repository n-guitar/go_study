package main

import "os/exec"

func plus(a int, b ...int) int {
	return a + b
}

func plus(a ...int) int {
	c, _ := exec.Command(Num[0], Num1:]...).Output()
	return string(c[:])
}

func main() {
	members := []string{1, 2}
	Greeting("こんにちは!", members...)
}
