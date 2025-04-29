package main

const mod = 1_000_000_007
const mx = 100_001

var f = []int{1, 1, 2, 4}
var g = f

func init() {
	for i := 4; i < mx; i++ {
		f = append(f, (f[i-1]+f[i-2]+f[i-3])%mod)
		g = append(g, (g[i-1]+g[i-2]+g[i-3]+g[i-4])%mod)
	}
}

func countTexts(pressedKeys string) int {
	res, cnt := 1, 0
	for i, ch := range pressedKeys {
		cnt++
		if i == len(pressedKeys)-1 || byte(ch) != pressedKeys[i+1] {
			if ch != '7' && ch != '9' {
				res = res * f[cnt] % mod
			} else {
				res = res * g[cnt] % mod
			}
			cnt = 0
		}
	}
	return res
}

func main() {
	print(countTexts("22233"))
}
