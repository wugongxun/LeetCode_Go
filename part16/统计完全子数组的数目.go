package main

func main() {
	print(countCompleteSubarrays([]int{1, 3, 1, 2, 2}))
}

func countCompleteSubarrays(nums []int) int {
	set := map[int]any{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	k := len(set)
	res, left := 0, 0
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
		for len(cnt) == k {
			cnt[nums[left]]--
			if cnt[nums[left]] == 0 {
				delete(cnt, nums[left])
			}
			left++
		}
		res += left
	}
	return res
}
