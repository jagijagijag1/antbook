package main

import "fmt"

var N, M int
var field []string

func dfsLakeCounting(x int, y int) {
	tmp_rune := []rune(field[x])
	tmp_rune[y] = '.'
	field[x] = string(tmp_rune)

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx, ny := x+dx, y+dy

			if 0 <= nx && nx < N &&
				0 <= ny && ny < M &&
				field[nx][ny] == 'W' {
				dfsLakeCounting(nx, ny)
			}
		}
	}
}

func solveLakeCounting(n int, m int, f []string) int {
	N, M = n, m
	field = f
	res := 0

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if field[i][j] == 'W' {
				dfsLakeCounting(i, j)
				res++
			}
		}
	}

	fmt.Println(res)
	return res
}
