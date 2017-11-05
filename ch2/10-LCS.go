package main

import "fmt"

var lcs_dp [][]int

func solveLCS(n, m int, s, t string) int {
	s, t = "0"+s, "0"+t
	initLCSDP()
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			//fmt.Println("i =", i, ", j =", j)
			if s[i] == t[j] {
				lcs_dp[i][j] = lcs_dp[i-1][j-1] + 1
				//fmt.Println("  (", i+1, ", ", j+1, ") is ", lcs_dp[i][j], " + 1")
			} else {
				lcs_dp[i][j] = max(lcs_dp[i-1][j], lcs_dp[i][j-1])
				//fmt.Println("  (", i+1, ", ", j+1, ") is ", lcs_dp[i][j+1], " or ", lcs_dp[i+1][j])
			}
			//printLCSDP(n+1, m+1, s, t)
		}
	}

	fmt.Println(lcs_dp[n][m])
	return lcs_dp[n][m]
}

func initLCSDP() {
	MAX_N, MAX_M := 1000, 1000
	lcs_dp = make([][]int, MAX_N+1)
	for i := 0; i < MAX_N; i++ {
		lcs_dp[i] = make([]int, MAX_M+1)
		for j := 0; j < MAX_M; j++ {
			lcs_dp[i][j] = 0
		}
	}
}

func printLCSDP(n, m int, s, t string) {
	fmt.Println(" ", t)
	for i := 0; i < n; i++ {
		fmt.Print(string(s[i]), " ")
		for j := 0; j < m; j++ {
			fmt.Print(lcs_dp[i][j], "")
		}
		fmt.Println()
	}
}
