package main

import "fmt"

func solveSarumansArmy(N int, R int, X []int) int {
	i, ans := 0, 0

	for i < N {
		s := X[i]
		for i < N && X[i] <= s+R {
			i++
		}

		p := X[i-1]
		for i < N && X[i] <= p+R {
			i++
		}

		ans++
	}

	fmt.Println(ans)
	return ans
}
