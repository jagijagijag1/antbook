package main

import "fmt"
import "container/list"

const INF = 100000000

type P struct {
	x, y interface{}
}

var maze []string
var sx, sy, gx, gy int

var dx, dy = []int{1, 0, -1, 0}, []int{0, 1, 0, -1}

func bfsMazeShortestPath(N int, M int) int {
	que := list.New()

	d := make([][]int, N)
	for i := range d {
		d[i] = make([]int, M)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			d[i][j] = INF

			if maze[i][j] == 'S' {
				sx, sy = i, j
			} else if maze[i][j] == 'G' {
				gx, gy = i, j
			}
		}
	}

	que.PushBack(P{sx, sy})
	d[sx][sy] = 0

	for que.Len() != 0 {
		p := que.Remove(que.Front())
		if p.(P).x.(int) == gx && p.(P).y.(int) == gy {
			break
		}

		for i := 0; i < 4; i++ {
			nx, ny := p.(P).x.(int)+dx[i], p.(P).y.(int)+dy[i]

			if 0 <= nx && nx < N &&
				0 <= ny && ny < M &&
				maze[nx][ny] != '#' &&
				d[nx][ny] == INF {
				que.PushBack(P{nx, ny})
				d[nx][ny] = d[p.(P).x.(int)][p.(P).y.(int)] + 1
			}
		}
	}

	return d[gx][gy]
}

func solveMazeShortestPath(N int, M int, a []string) int {
	maze = a
	res := bfsMazeShortestPath(N, M)

	fmt.Println(res)
	return res
}
