package main

import "fmt"

func solveFenceRepair(N int, L []int) int {
	ans := 0

	for N > 1 {
		mii1, mii2 := 0, 1
		if L[mii1] > L[mii2] {
			L[mii1], L[mii2] = L[mii2], L[mii1]
		}

		for i := 2; i < N; i++ {
			if L[i] < L[mii1] {
				mii2, mii1 = mii1, i
			} else if L[i] < L[mii2] {
				mii2 = i
			}
		}

		t := L[mii1] + L[mii2]
		ans += t

		if mii1 == N-1 {
			mii1, mii2 = mii2, mii1
		}
		L[mii1], L[mii2] = t, L[N-1]
		N--
	}

	fmt.Println(ans)
	return ans
}
