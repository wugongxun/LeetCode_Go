package main

import (
	"math"
	"slices"
	"strconv"
)

func main() {
	print(countGoodIntegers(5, 6))
}

func countGoodIntegers(n int, k int) (res int64) {
	fac := make([]int, n+1)
	fac[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i
	}
	vis := map[string]bool{}
	base := int(math.Pow10((n - 1) / 2))
	for i := base; i < 10*base; i++ {
		x, t := i, i
		if n%2 != 0 {
			t /= 10
		}
		for ; t > 0; t /= 10 {
			x = x*10 + t%10
		}
		if x%k != 0 {
			continue
		}
		bs := []byte(strconv.Itoa(x))
		slices.Sort(bs)
		sortedStr := string(bs)
		if vis[sortedStr] {
			continue
		}
		vis[sortedStr] = true

		cnt := [10]int{}
		for _, c := range bs {
			cnt[c-'0']++
		}
		tmp := (n - cnt[0]) * fac[n-1]
		for _, c := range cnt {
			tmp /= fac[c]
		}
		res += int64(tmp)
	}
	return
}
