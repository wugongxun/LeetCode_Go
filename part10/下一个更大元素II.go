package main

import "fmt"

func main() {
	fmt.Printf("%v", nextGreaterElements([]int{1, 2, 3, 2, 1}))
}

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	var stack []int
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	for i := 0; i < 2*n; i++ {
		cur := nums[i%n]
		for len(stack) > 0 && cur > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = cur
			stack = stack[:len(stack)-1]
		}
		if i < n {
			stack = append(stack, i)
		}
	}
	return res
}
