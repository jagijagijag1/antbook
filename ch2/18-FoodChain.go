package main

/* union-find implementation */
var par, rank []int

func init_tree(n int) {
	par = make([]int, n)
	rank = make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = i
		rank[i] = 0
	}
}

func find(x int) int {
	if par[x] == x {
		return x
	} else {
		par[x] = find(par[x])
		return par[x]
	}
}

func unite(x, y int) {
	x, y = find(x), find(y)
	if x == y {
		return
	}

	if rank[x] < rank[y] {
		par[x] = y
	} else {
		par[y] = x
		if rank[x] == rank[y] {
			rank[x]++
		}
	}
}

func same(x, y int) bool {
	return find(x) == find(y)
}

/* main */
func solveFoodChain(N, K int, T, X, Y []int) int {
	init_tree(N * 3)

	ans := 0
	for i := 0; i < K; i++ {
		t, x, y := T[i], X[i]-1, Y[i]-1

		if x < 0 || N <= x || y < 0 || N <= y {
			ans++
			continue
		}

		if t == 1 {
			if same(x, y+N) || same(x, y+2*N) {
				ans++
			} else {
				unite(x, y)
				unite(x+N, y+N)
				unite(x+N*2, y+N*2)
			}
		} else {
			if same(x, y) || same(x, y+2*N) {
				ans++
			} else {
				unite(x, y+N)
				unite(x+N, y+2*N)
				unite(x+2*N, y)
			}
		}
	}

	return ans
}
