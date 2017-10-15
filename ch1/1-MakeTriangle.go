package main

import "fmt"
import "math"

func solveMakeTriangle(n int, a []int) int {
	ans := 0

	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				len := a[i] + a[j] + a[k]
				ma := int(math.Max(float64(a[i]), math.Max(float64(a[j]), float64(a[k]))))
				rest := len - ma

				if ma < rest {
					ans = int(math.Max(float64(ans), float64(len)))
				}
			}
		}
	}

	fmt.Println(ans)
	return ans
}
