package main

func main() {

}

func rob(nums []int) int {
	d0, d1 := 0, 0
	for _, num := range nums {
		d := max(d0+num, d1)
		d0 = d1
		d1 = d
	}
	return max(d0, d1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
