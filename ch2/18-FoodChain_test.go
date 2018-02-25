package main

import "testing"

func TestFoodChain(t *testing.T) {
	N, K := 100, 7
	T := []int{1, 2, 2, 2, 1, 2, 1}
	X := []int{101, 1, 2, 3, 1, 3, 5}
	Y := []int{1, 2, 3, 3, 3, 1, 5}

	actual := solveFoodChain(N, K, T, X, Y)
	expected := 3
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
