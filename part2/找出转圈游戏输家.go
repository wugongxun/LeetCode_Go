package main

import "fmt"

func main() {
	fmt.Printf("%v", circularGameLosers(5, 2))
}

func circularGameLosers(n int, k int) []int {
	visit := make([]bool, n)
	j := 0
	for i := k; !visit[j]; i += k {
		visit[j] = true
		j = (j + i) % n
	}
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !visit[i] {
			res = append(res, i+1)
		}
	}
	return res
}
