package main

import "testing"

func TestPartialSum1(t *testing.T) {
	n := 4
	a := []int{1, 2, 4, 7}
	k := 13

	actual := solvePartialSum(n, a, k)
	expected := true
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}

func TestPartialSum2(t *testing.T) {
	n := 4
	a := []int{1, 2, 4, 7}
	k := 15

	actual := solvePartialSum(n, a, k)
	expected := false
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
