package main

import "testing"

func TestUnboundedKnapsack(t *testing.T) {
	n := 3
	w, v := []int{3, 4, 2}, []int{4, 5, 3}
	W := 7

	actual := solveUnboundedKnapsack(n, w, v, W)
	expected := 10
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}

	actual = solveUnboundedKnapsack1DArray(n, w, v, W)
	expected = 10
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
