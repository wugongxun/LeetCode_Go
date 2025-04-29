package main

import "slices"

func countSubarrays(nums []int, k int) (res int64) {
	mx := slices.Max(nums)
	i0, cnt, n := 0, 0, len(nums)
	for i, x := range nums {
		if x == mx {
			cnt++
		}
		for cnt >= k {
			res += int64(n - i)
			if nums[i0] == mx {
				cnt--
			}
			i0++
		}
	}
	return
}

func main() {
	print(countSubarrays([]int{1, 3, 2, 3, 3}, 2))
}
