package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v", wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

func wordBreak(s string, wordDict []string) []string {
	n := len(s)
	dp := make([][]string, n+1)
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			if len(dp[j]) == 0 && j != 0 {
				continue
			}
			t := s[j:i]
			for _, word := range wordDict {
				if word == t {
					if j == 0 {
						dp[i] = append(dp[i], t)
					} else {
						for _, prev := range dp[j] {
							dp[i] = append(dp[i], prev+" "+t)
						}
					}
					break
				}
			}
		}
	}
	return dp[n]
}
