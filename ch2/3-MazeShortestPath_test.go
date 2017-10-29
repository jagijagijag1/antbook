package main

import "testing"

func TestMazeShortestPath(t *testing.T) {
	N, M := 10, 10
	maze := []string{
		"#S######.#",
		"......#..#",
		".#.##.##.#",
		".#........",
		"##.##.####",
		"....#....#",
		".#######.#",
		"....#.....",
		".####.###.",
		"....#...G#",
	}

	actual := solveMazeShortestPath(N, M, maze)
	expected := 22
	if actual != expected {
		t.Errorf("expected %v, but got %v\n", expected, actual)
	}
}
