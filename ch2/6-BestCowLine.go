package main

import "fmt"

func solveBestCowLine(N int, S string) string {
	res := ""
	a, b := 0, N-1

	for a <= b {
		left := false

		for i := 0; a+i <= b; i++ {
			if S[a+i] < S[b-i] {
				left = true
				break
			} else if S[a+i] > S[b-i] {
				left = false
				break
			}
		}

		if left {
			res += string(S[a])
			a++
		} else {
			res += string(S[b])
			b--
		}
	}

	fmt.Println(res)
	return res
}
