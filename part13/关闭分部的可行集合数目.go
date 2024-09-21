package main

import "math"

func numberOfSets(n int, maxDistance int, roads [][]int) (res int) {
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, n)
		for j := range adj[i] {
			adj[i][j] = math.MaxInt / 2
		}
	}
	for _, road := range roads {
		x, y, w := road[0], road[1], road[2]
		adj[x][y] = min(adj[x][y], w)
		adj[y][x] = min(adj[y][x], w)
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
label:
	for s := 0; s < 1<<n; s++ {
		for i, row := range adj {
			if s>>i&1 > 0 {
				copy(f[i], row)
			}
		}

		for k := range f {
			if (s >> k & 1) == 0 {
				continue
			}
			for i := range f {
				if (s >> i & 1) == 0 {
					continue
				}
				for j := range f {
					f[i][j] = min(f[i][j], f[i][k]+f[k][j])
				}
			}
		}

		for i, di := range f {
			if (s >> i & 1) == 0 {
				continue
			}
			for j, dj := range di[:i] {
				if s>>j&1 > 0 && dj > maxDistance {
					continue label
				}
			}
		}
		res++
	}
	return
}
