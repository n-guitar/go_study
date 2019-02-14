package main

import (
	"testing"
)

func TestArith(t *testing.T) {
	//期待する結果
	expplus := 12
	expminus := 8
	expmultiple := 20
	expdivide := 5
	//実際の結果
	resplus, resminus, resmultiple, resdivide := arithmetic(10, 2)
	//結果の比較
	if expplus != resplus {
		t.Errorf("!!ERROR!!You got %v for \"+\". %v is expected.\n", resplus, expplus)
	}
	if expminus != resminus {
		t.Errorf("!!ERROR!!You got %v for \"-\". %v is expected.\n", resminus, expminus)
	}
	if expmultiple != resmultiple {
		t.Errorf("!!ERROR!!You got %v for \"×\". %v is expected.\n", resmultiple, expmultiple)
	}
	if expdivide != resdivide {
		t.Errorf("!!ERROR!!You got %v for \"÷\". %v is expected.\n", resdivide, expdivide)
	}
}
