package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestMul(t *testing.T) {
	var total int64 = Mul(5, 5)
	var expected int64 = 25
	if total != expected {
		t.Errorf("Mul was incorrect, got: %d, want: %d.", total, expected)
	}
}
