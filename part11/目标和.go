package main

func main() {
	print(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, target int) int {
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 != 0 {
		return 0
	}
	target /= 2
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums {
		for s := target; s >= num; s-- {
			dp[s] += dp[s-num]
		}
	}
	return dp[target]
}
