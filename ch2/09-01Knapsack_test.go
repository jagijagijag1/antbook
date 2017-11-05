package main

import "testing"

func Test01Knapsack(t *testing.T) {
	n := 4
	w, v := []int{2, 1, 3, 2}, []int{3, 2, 4, 2}
	W := 5

	actual := solve01Knapsack(n, w, v, W)
	expected := 7
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
