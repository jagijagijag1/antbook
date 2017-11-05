package main

import "testing"

func TestGreedyCoin(t *testing.T) {
	C := []int{3, 2, 1, 3, 0, 2}
	A := 620

	actual := solveGreedyCoin(C, A)
	expected := 6
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
