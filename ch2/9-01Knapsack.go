package main

import "fmt"

const MAX_N, MAX_W = 100, 100

var dp [][]int

func solve01Knapsack(n int, w []int, v []int, W int) int {
	initDP()

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= W; j++ {
			if j < w[i] {
				dp[i][j] = dp[i+1][j]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i+1][j-w[i]]+v[i])
			}
		}
	}

	fmt.Println(dp[0][W])
	return dp[0][W]
}

func initDP() {
	dp = make([][]int, MAX_N+1)
	for i := 0; i < MAX_N; i++ {
		dp[i] = make([]int, MAX_W+1)
		for j := 0; j < MAX_W; j++ {
			dp[i][j] = 0
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
