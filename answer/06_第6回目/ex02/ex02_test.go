package main

import (
	"testing"
)

// 参照：https://qiita.com/high5/items/4e2580241039c950e1c4
func Calc(a, b int) (int, int, int, float32) {
	plus := a + b
	minus := a - b
	multi := a * b
	divi := float32(a) / float32(b)
	return plus, minus, multi, divi
}

func TestCalc(t *testing.T) {

	a, b := 1, 2

	Plus_want := 3
	Minus_want := -1
	Multi_want := 2
	Divi_want := float32(0.5)

	plus, minus, multi, divi := Calc(a, b)

	if plus != Plus_want {
		t.Errorf("%d is not Plus ans", plus)
	} else if minus != Minus_want {
		t.Errorf("%d is not Minus ans", minus)
	} else if multi != Multi_want {
		t.Errorf("%d is not Multi ans", multi)
	} else if divi != Divi_want {
		t.Errorf("%f is not Divi ans", divi)
	}
}
