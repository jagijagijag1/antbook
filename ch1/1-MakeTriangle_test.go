package main

import "testing"

func TestMakeTriangle1(t *testing.T) {
	n := 5
	a := []int{2, 3, 4, 5, 10}

	actual := solveMakeTriangle(n, a)
	expected := 12
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}

func TestMakeTriangle2(t *testing.T) {
	n := 4
	a := []int{4, 5, 10, 20}

	actual := solveMakeTriangle(n, a)
	expected := 0
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
