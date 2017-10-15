package main

import "fmt"
import "sort"

const max_n = 50

var kk = make([]int, max_n*max_n, max_n*max_n)

func solveSumOfFourLots(n int, m int, k []int) bool {
	for c := 0; c < n; c++ {
		for d := 0; d < n; d++ {
			kk[c*n+d] = k[c] + k[d]
		}
	}

	sort.Ints(kk)

	f := false
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			if binary_search(m-k[a]-k[b], n) {
				f = true
			}
		}
	}

	fmt.Println(f)
	return f
}

func binary_search(x int, n int) bool {
	l, r := 0, n*n

	for r-l >= 1 {
		i := (l + r) / 2
		if kk[i] == x {
			return true
		} else if kk[i] < x {
			l = i + 1
		} else {
			r = i
		}
	}

	return false
}
