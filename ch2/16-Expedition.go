package main

import "fmt"
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // larger int has higher priority
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solveExpedition(N, L, P int, A, B []int) int {
	A = append(A, L)
	B = append(B, 0)
	N++

	que := &IntHeap{}
	heap.Init(que)

	ans, pos, tank := 0, 0, P

	for i := 0; i < N; i++ {
		d := A[i] - pos

		for tank-d < 0 {
			if que.Len() == 0 {
				return -1
			}

			tank += heap.Pop(que).(int)
			ans++
		}

		tank -= d
		pos = A[i]
		heap.Push(que, B[i])
	}

	fmt.Println(ans)
	return ans
}
