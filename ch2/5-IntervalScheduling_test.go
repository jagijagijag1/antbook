package main

import "testing"

func TestIntervalScheduling(t *testing.T) {
	n := 5
	S := []int{1, 2, 4, 6, 8}
	T := []int{3, 5, 7, 9, 10}

	actual := solveIntervalScheduling(n, S, T)
	expected := 3
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
