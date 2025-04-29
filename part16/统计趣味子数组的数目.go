package main

func countInterestingSubarrays(nums []int, modulo int, k int) (res int64) {
	n := len(nums)
	prefix := make([]int, n+1)
	for i, num := range nums {
		prefix[i+1] = prefix[i]
		if num%modulo == k {
			prefix[i+1]++
		}
	}
	cnt := make([]int, min(n+1, modulo))
	for _, x := range prefix {
		if x >= k {
			res += int64(cnt[(x-k)%modulo])
		}
		cnt[x%modulo]++
	}
	return
}

func main() {
	println(countInterestingSubarrays([]int{3, 2, 4}, 2, 1))
	println(countInterestingSubarrays([]int{3, 1, 9, 6}, 3, 0))
}
