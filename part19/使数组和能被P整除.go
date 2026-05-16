package main

// 实际题目要求是找出在前缀和（ps）中，找到和nums的元素和（s）同余
// 即：ps[r] − ps[l] ≡ s(mod p)
// 移项可得：(ps[r] mod p - s mod p + p) mod p = ps[l] mod p
func minSubarray(nums []int, p int) int {
	n := len(nums)
	// 前缀和，但是提前 mod p -> 相当于 ps[r] mod p
	ps := make([]int, n+1)
	for i, x := range nums {
		ps[i+1] = (ps[i] + x) % p
	}
	// 这个s就是s mod p，因为提前mod p了
	s := ps[n]
	// 不用删除子数组了，直接返回
	if s == 0 {
		return 0
	}
	res := n
	// last 用于记录最近的 ps[l] mod p 出现的下标
	last := map[int]int{}
	for i, x := range ps {
		last[x] = i
		if j, ok := last[(x-s+p)%p]; ok {
			res = min(res, i-j)
		}
	}
	if res < n {
		return res
	}
	return -1
}

func main() {
	println(minSubarray([]int{3, 1, 4, 2}, 6))
}
