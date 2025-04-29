package main

func countSubarrays(nums []int, minK int, maxK int) (res int64) {
	minI, maxI, i0 := -1, -1, -1
	for i, num := range nums {
		if num == minK {
			minI = i
		}
		if num == maxK {
			maxI = i
		}
		if num < minK || num > maxK {
			i0 = i
		}
		res += int64(max(min(minI, maxI)-i0, 0))
	}
	return
}

func main() {
	println(countSubarrays([]int{1, 3, 5, 2, 1, 5, 3}, 1, 5))
	//println(countSubarrays([]int{1, 1, 1, 1}, 1, 1))
}
