package main

import "fmt"

func solveUnboundedKnapsack(n int, w, v []int, W int) int {
	// init dp matrix
	MAX_N, MAX_W := 100, 100
	dp := make([][]int, MAX_N+1)
	for i := 0; i < MAX_N; i++ {
		dp[i] = make([]int, MAX_W+1)
		for j := 0; j < MAX_W; j++ {
			dp[i][j] = 0
		}
	}

	// main
	for i := 0; i < n; i++ {
		for j := 0; j <= W; j++ {
			if j < w[i] {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = max(dp[i][j], dp[i+1][j-w[i]]+v[i])
			}
		}
	}

	fmt.Println(dp[n][W])
	return dp[n][W]
}

func solveUnboundedKnapsack1DArray(n int, w, v []int, W int) int {
	// init dp matrix
	MAX_W := 100
	dp := make([]int, MAX_W+1)
	for i := 0; i < MAX_W; i++ {
		dp[i] = 0
	}

	// main
	for i := 0; i < n; i++ {
		for j := w[i]; j <= W; j++ {
			dp[j] = max(dp[j], dp[j-w[i]]+v[i])
			//fmt.Println(i, ",", j, ":", dp[0:10])
		}
	}

	fmt.Println(dp[W])
	return dp[W]
}
