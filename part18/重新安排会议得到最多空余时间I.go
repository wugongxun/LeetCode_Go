package main

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	free := make([]int, n+1)
	free[0] = startTime[0]
	for i := 1; i < n; i++ {
		free[i] = startTime[i] - endTime[i-1]
	}
	free[n] = eventTime - endTime[n-1]
	res, s := 0, 0
	for i, f := range free {
		s += f
		if i < k {
			continue
		}
		res = max(res, s)
		s -= free[i-k]
	}
	return res
}

func main() {
	println(maxFreeTime(10, 1, []int{0, 2, 9}, []int{1, 4, 10}))
}
