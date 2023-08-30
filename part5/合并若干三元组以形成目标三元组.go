package main

import "fmt"

func main() {
	fmt.Printf("%v", mergeTriplets([][]int{
		{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3},
	}, []int{5, 5, 5}))
}

func mergeTriplets(triplets [][]int, target []int) bool {
	a, b, c := 0, 0, 0
	for _, triplet := range triplets {
		if triplet[0] <= target[0] && triplet[1] <= target[1] && triplet[2] <= target[2] {
			a, b, c = max(a, triplet[0]), max(b, triplet[1]), max(c, triplet[2])
		}
	}
	return a == target[0] && b == target[1] && c == target[2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
