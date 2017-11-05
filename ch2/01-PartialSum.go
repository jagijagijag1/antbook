package main

func dfs(n int, i int, sum int, a []int, k int) bool {
	if i == n {
		return sum == k
	}

	if dfs(n, i+1, sum, a, k) {
		return true
	}

	if dfs(n, i+1, sum+a[i], a, k) {
		return true
	}

	return false
}

func solvePartialSum(n int, a []int, k int) bool {
	if dfs(n, 0, 0, a, k) {
		return true
	} else {
		return false
	}
}
