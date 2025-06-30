package main

func colorTheGrid(m int, n int) int {
	const mod = 1_000_000_007

	pow3 := make([]int, m)
	pow3[0] = 1
	for i := 1; i < m; i++ {
		pow3[i] = pow3[i-1] * 3
	}

	var valid []int
next:
	for color := range pow3[m-1] * 3 {
		for i := range m - 1 {
			if color/pow3[i+1]%3 == color/pow3[i]%3 {
				continue next
			}
		}
		valid = append(valid, color)
	}

	nv := len(valid)
	nxt := make([][]int, nv)
	for i, c1 := range valid {
	next2:
		for j, c2 := range valid {
			for _, p := range pow3 {
				if c1/p%3 == c2/p%3 {
					continue next2
				}
			}
			nxt[i] = append(nxt[i], j)
		}
	}

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, nv)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i == 0 {
			return 1
		}
		p := &cache[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		for _, k := range nxt[j] {
			res += dfs(i-1, k)
		}
		return res % mod
	}
	res := 0
	for j := range nv {
		res += dfs(n-1, j)
	}
	return res % mod
}

func main() {
	println(colorTheGrid(5, 5))
}
