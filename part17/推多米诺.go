package main

func pushDominoes(dominoes string) string {
	s := []byte("L" + dominoes + "R")
	pre := 0
	for i, ch := range s {
		if ch == '.' {
			continue
		}
		if ch == s[pre] {
			fill(s[pre:i], ch)
		} else if ch == 'L' {
			fill(s[pre:pre+(i-pre+1)/2], 'R')
			fill(s[pre+(i-pre)/2+1:i], 'L')
		}
		pre = i
	}
	return string(s[1 : len(s)-1])
}

func fill(s []byte, ch byte) {
	for i := range s {
		s[i] = ch
	}
}

func main() {
	print(pushDominoes(".L.R...LR..L.."))
}
