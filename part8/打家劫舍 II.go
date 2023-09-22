package main

func main() {

}

//func rob(nums []int) int {
//	n := len(nums)
//	if n == 1 {
//		return nums[0]
//	}
//	if n == 2 {
//		return max(nums[0], nums[1])
//	}
//	return max(_rob(nums, 0, n-1), _rob(nums, 1, n))
//}

func _rob(nums []int, start, end int) int {
	d0, d1 := 0, 0
	for i := start; i < end; i++ {
		d := max(d0+nums[i], d1)
		d0 = d1
		d1 = d
	}
	return max(d0, d1)
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
