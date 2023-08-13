package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 0, 0, 0}
	merge(arr, 3, []int{2, 5, 6}, 3)
	fmt.Printf("%v", arr)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	index := n + m - 1
	m--
	n--
	for n >= 0 {
		if m >= 0 && nums1[m] >= nums2[n] {
			nums1[index] = nums1[m]
			m--
		} else {
			nums1[index] = nums2[n]
			n--
		}
		index--
	}
}

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	res := make([]int, 0, m+n)
//	i, j := 0, 0
//	for {
//		if i == m {
//			res = append(res, nums2[j:]...)
//			break
//		}
//		if j == n {
//			res = append(res, nums1[i:]...)
//			break
//		}
//		if nums1[i] < nums2[j] {
//			res = append(res, nums1[i])
//			i++
//		} else {
//			res = append(res, nums2[j])
//			j++
//		}
//	}
//	copy(nums1, res)
//}
