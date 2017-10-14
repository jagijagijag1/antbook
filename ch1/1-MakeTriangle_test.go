package triangle

import "testing"

func Test1(t *testing.T) {
	n := 5
	a := []int{2, 3, 4, 5, 10}

	actual := solve(n, a)
	expected := 12
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}

func Test2(t *testing.T) {
	n := 4
	a := []int{4, 5, 10, 20}

	actual := solve(n, a)
	expected := 0
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
