package main

import "testing"

func TestLakeCounting1(t *testing.T) {
	N, M := 10, 12
	field := []string{
		"W.........WW.",
		".WWW......WWW",
		"....WW....WW.",
		"..........WW.",
		"..........W..",
		"..W.......W..",
		".W.W......WW.",
		"W.W.W......W.",
		".W.W.......W.",
		"..W........W.",
	}

	actual := solveLakeCounting(N, M, field)
	expected := 3
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
