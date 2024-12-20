package main

func minAnagramLength(s string) int {
	n := len(s)
	cntAll := [26]int{}
	for _, c := range s {
		cntAll[c-'a']++
	}
	g := 0
	for _, c := range cntAll {
		g = gcd(g, c)
	}
label:
	for times := g; times > 1; times-- {
		if g%times != 0 {
			continue
		}
		k := n / times
		cnt1 := [26]int{}
		for _, c := range s[:k] {
			cnt1[c-'a']++
		}
		for i := k; i < n; i += k {
			cnt2 := [26]int{}
			for _, c := range s[i : i+k] {
				cnt2[c-'a']++
			}
			if cnt1 != cnt2 {
				continue label
			}
		}
		return k
	}
	return n
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
