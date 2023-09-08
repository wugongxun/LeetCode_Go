package main

import "fmt"

func main() {
	fmt.Printf("%v", candy([]int{1, 2, 3}))
	//fmt.Printf("%v", candy([]int{1, 3, 4, 5, 2}))
	//fmt.Printf("%v", candy([]int{1, 2, 2}))
}

func candy(ratings []int) (res int) {
	n := len(ratings)
	left := make([]int, n)
	for i, rating := range ratings {
		if i > 0 && rating > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		res += max(left[i], right)
	}
	return
}

//func candy(ratings []int) int {
//	n := len(ratings)
//	if n == 1 {
//		return 1
//	}
//	if n == 2 {
//		if ratings[0] == ratings[1] {
//			return 2
//		} else {
//			return 3
//		}
//	}
//	lows := make([]int, 0)
//	if ratings[1] >= ratings[0] {
//		lows = append(lows, 0)
//	}
//	for i := 1; i < n-1; i++ {
//		if ratings[i] <= ratings[i-1] && ratings[i] <= ratings[i+1] {
//			lows = append(lows, i)
//		}
//	}
//	if ratings[n-2] >= ratings[n-1] {
//		lows = append(lows, n-1)
//	}
//	cnt := make([]int, n)
//	for _, low := range lows {
//		cnt[low] = 1
//		i := low
//		for i+1 < n && ratings[i] < ratings[i+1] {
//			cnt[i+1] = max(cnt[i+1], cnt[i]+1)
//			i++
//		}
//		i = low
//		for i-1 >= 0 && ratings[i] < ratings[i-1] {
//			cnt[i-1] = max(cnt[i-1], cnt[i]+1)
//			i--
//		}
//	}
//	res := 0
//	for _, v := range cnt {
//		res += v
//	}
//	return res
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
