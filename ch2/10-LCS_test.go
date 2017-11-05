package main

import "testing"

func TestLCS(t *testing.T) {
	n, m := 4, 4
	ss, st := "abcd", "becd"

	actual := solveLCS(n, m, ss, st)
	expected := 3
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
