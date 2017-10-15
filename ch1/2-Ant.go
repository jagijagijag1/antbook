package main

import "fmt"
import "math"

func solveAnt(L int, n int, x []int) (int, int) {
	min := 0.0
	max := 0.0

	for i := 0; i < n; i++ {
		min = math.Max(min, math.Min(float64(x[i]), float64(L-x[i])))
	}

	for i := 0; i < n; i++ {
		fmt.Println("\t", float64(x[i]))
		max = math.Max(max, math.Max(float64(x[i]), float64(L-x[i])))
	}

	fmt.Println("min:", int(min), ", max:", int(max))
	return int(min), int(max)
}
