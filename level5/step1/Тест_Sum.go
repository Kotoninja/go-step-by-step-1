package main

import "testing"

func TestSum(t *testing.T) {
	var a, b int = 2, 5
	sumAnswer := Sum(a + b)
	if a+b != sumAnswer {
		t.Errorf("Test sum failed")
	}
}
