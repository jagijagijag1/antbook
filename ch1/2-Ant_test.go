package main

import "testing"

func TestAnt(t *testing.T) {
	L := 10
	n := 3
	x := []int{2, 6, 7}

	actualMin, actualMax := solveAnt(L, n, x)
	expectedMin, expectedMax := 4, 8
	if actualMin != expectedMin || actualMax != expectedMax {
		t.Errorf("expected <%v, %v>, but got <%v, %v>\n", expectedMin, expectedMax, actualMin, actualMax)
	}
}
