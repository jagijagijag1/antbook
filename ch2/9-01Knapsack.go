package main

import "fmt"

const MAX_N, MAX_W = 100, 100

var dp [][]int
var n, W int
var w, v []int

func rec(i int, j int) int {
	if dp[i][j] >= 0 {
		return dp[i][j]
	}

	var res int
	if i == n {
		res = 0
	} else if j < w[i] {
		res = rec(i+1, j)
	} else {
		res = max(rec(i+1, j), rec(i+1, j-w[i])+v[i])
	}

	dp[i][j] = res
	return res
}

func solve01Knapsack(n_in int, w_in []int, v_in []int, W_in int) int {
	n, w, v, W = n_in, w_in, v_in, W_in
	initDP()

	ans := rec(0, W)
	fmt.Println(ans)
	return ans
}

func initDP() {
	dp = make([][]int, MAX_N+1)
	for i := 0; i < MAX_N; i++ {
		dp[i] = make([]int, MAX_W+1)
		for j := 0; j < MAX_W; j++ {
			dp[i][j] = -1
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
