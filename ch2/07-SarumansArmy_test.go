package main

import "testing"

func TestSarumansArmy(t *testing.T) {
	N, R := 6, 10
	X := []int{1, 7, 15, 20, 30, 50}

	actual := solveSarumansArmy(N, R, X)
	expected := 3
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
