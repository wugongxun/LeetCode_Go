package main

import "math"

func robotWithString(s string) string {
	n := len(s)
	sufMin := make([]byte, n+1)
	sufMin[n] = math.MaxUint8
	for i := n - 1; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], s[i])
	}
	res := make([]byte, 0, n)
	var st []byte
	for i, ch := range s {
		st = append(st, byte(ch))
		for len(st) > 0 && st[len(st)-1] <= sufMin[i+1] {
			res = append(res, st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return string(res)
}

func main() {
	println(robotWithString("zza"))
}
