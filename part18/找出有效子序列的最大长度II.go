package main

func maximumLength(nums []int, k int) (res int) {
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	for _, x := range nums {
		x %= k
		for y, fxy := range dp[x] {
			dp[y][x] = fxy + 1
			res = max(res, dp[y][x])
		}
	}
	return
}

func main() {
	println(maximumLength([]int{1, 2, 3, 4, 5}, 2))
}
