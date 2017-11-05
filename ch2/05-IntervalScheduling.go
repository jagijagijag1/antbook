package main

import "fmt"
import "sort"

type Interval struct {
	end, start int
}

func solveIntervalScheduling(n int, S []int, T []int) int {
	itv := make([]Interval, n)
	for i := 0; i < n; i++ {
		itv[i] = Interval{T[i], S[i]}
	}
	sort.Slice(itv, func(i, j int) bool {
		return itv[i].end < itv[j].end
	})

	ans, last_t := 0, 0
	for i := 0; i < n; i++ {
		if last_t < itv[i].start {
			ans++
			last_t = itv[i].end
		}
	}

	fmt.Println(ans)
	return ans
}
