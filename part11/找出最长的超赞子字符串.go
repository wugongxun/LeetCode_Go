package main

func main() {
	print(longestAwesome("3242415"))
}

func longestAwesome(s string) (res int) {
	n := len(s)
	pos := make([]int, 1<<10)
	for i := range pos {
		pos[i] = n
	}
	pos[0] = -1
	pre := 0
	for i, c := range s {
		pre ^= 1 << (c - '0')
		for d := 0; d < 10; d++ {
			res = max(res, i-pos[pre^(1<<d)])
		}
		res = max(res, i-pos[pre])
		if pos[pre] == n {
			pos[pre] = i
		}
	}
	return res
}
