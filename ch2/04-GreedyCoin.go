package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solveGreedyCoin(C []int, A int) int {
	V := []int{1, 5, 10, 50, 100, 500}
	ans := 0

	for i := 5; i >= 0; i-- {
		t := min(A/V[i], C[i])
		A -= t * V[i]
		ans += t
	}

	fmt.Println(ans)
	return ans
}
