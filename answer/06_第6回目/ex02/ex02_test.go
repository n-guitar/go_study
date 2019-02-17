package main

import "testing"

func TestDuck(t *testing.T) {
	duck := &Duck{"tarou"}
	actual := duck.name
	expected := "tarou"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
