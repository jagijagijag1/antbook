package main

import "testing"

func TestSumOfFourLots1(t *testing.T) {
	n := 3
	m := 10
	k := []int{1, 3, 5}

	actual := solveSumOfFourLots(n, m, k)
	expected := true
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}

func TestSumOfFourLots2(t *testing.T) {
	n := 3
	m := 9
	k := []int{1, 3, 5}

	actual := solveSumOfFourLots(n, m, k)
	expected := false
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
