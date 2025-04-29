package main

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	res, i0, i1 := 0, -1, -1
	for i, num := range nums {
		if num > right {
			i0, i1 = i, i
		}
		if num >= left {
			i1 = i
		}
		res += i1 - i0
	}
	return res
}

func main() {
	println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))
	println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8))
	println(numSubarrayBoundedMax([]int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52}, 32, 69))
}
