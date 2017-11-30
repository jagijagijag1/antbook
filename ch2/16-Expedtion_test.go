package main

import "testing"

func TestExpedition(t *testing.T) {
	N, L, P := 4, 25, 10
	A, B := []int{10, 14, 20, 21}, []int{10, 5, 2, 4}

	actual := solveExpedition(N, L, P, A, B)
	expected := 2
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
