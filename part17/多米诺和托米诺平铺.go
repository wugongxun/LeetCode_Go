package main

func numTilings(n int) int {
	const mod = 1_000_000_007
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1]*2 + dp[i-3]) % mod
	}
	return dp[n]
}

//func numTilings(n int) int {
//	const mod = 1_000_000_007
//	cache := make([][]int, n)
//	for i := range cache {
//		cache[i] = make([]int, 3)
//		for j := range cache[i] {
//			cache[i][j] = -1
//		}
//	}
//	var dfs func(i, stat int) int
//	dfs = func(i, stat int) int {
//		if i == n {
//			if stat == 0 {
//				return 1
//			}
//			return 0
//		}
//		if cache[i][stat] != -1 {
//			return cache[i][stat]
//		}
//		res := 0
//		if stat == 0 {
//			res += dfs(i+1, 0) + dfs(i+1, 1) + dfs(i+1, 2)
//			if i+2 <= n {
//				res += dfs(i+2, 0)
//			}
//		} else {
//			res += dfs(i+1, stat^3)
//			if i+2 <= n {
//				res += dfs(i+2, 0)
//			}
//		}
//		cache[i][stat] = res % mod
//		return cache[i][stat]
//	}
//	return dfs(0, 0)
//}

func main() {
	println(numTilings(4))
}
