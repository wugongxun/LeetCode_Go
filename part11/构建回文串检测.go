package main

func main() {

}

func canMakePaliQueries(s string, queries [][]int) (res []bool) {
	n := len(s)
	pre := make([]int, n+1)
	for i := range n {
		pre[i+1] = pre[i] ^ 1<<(s[i]-'a')
	}
	for _, query := range queries {
		l, r, k := query[0], query[1], query[2]
		x := pre[l] ^ pre[r+1]
		cnt := 0
		for x != 0 {
			x = x & (x - 1)
			cnt++
		}
		res = append(res, cnt/2 <= k)
	}
	return
}
