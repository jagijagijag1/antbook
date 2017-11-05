package main

import "testing"

func TestBestCowLine(t *testing.T) {
	N := 6
	S := "ACDBCB"

	actual := solveBestCowLine(N, S)
	expected := "ABCBCD"
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
