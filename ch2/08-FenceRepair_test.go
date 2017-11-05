package main

import "testing"

func TestFenceRepair(t *testing.T) {
	N := 3
	L := []int{8, 5, 8}

	actual := solveFenceRepair(N, L)
	expected := 34
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
