package main

import (
	"testing"
)

func Test_add(t *testing.T) {
	if add(4, 2) != 6 {
		t.Errorf("invalid.")
	}
}

func Test_minus(t *testing.T) {
	if minus(4, 2) != 2 {
		t.Errorf("invalid.")
	}
}

func Test_multiply(t *testing.T) {
	if multiply(4, 2) != 8 {
		t.Errorf("invalid.")
	}
}

func Test_divide(t *testing.T) {
	if divide(4, 2) != 2 {
		t.Errorf("invalid.")
	}
}
