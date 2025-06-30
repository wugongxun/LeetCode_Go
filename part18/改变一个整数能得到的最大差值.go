package main

import (
	"strconv"
	"strings"
)

func maxDiff(num int) int {
	s := strconv.Itoa(num)
	mx := num
	for i, ch := range s {
		if ch != '9' {
			mx, _ = strconv.Atoi(strings.ReplaceAll(s, s[i:i+1], "9"))
			break
		}
	}
	mn := num
	for i, ch := range s {
		if i == 0 && ch != '1' {
			mn, _ = strconv.Atoi(strings.ReplaceAll(s, s[0:1], "1"))
			break
		}
		if ch > '1' {
			mn, _ = strconv.Atoi(strings.ReplaceAll(s, s[i:i+1], "0"))
			break
		}
	}
	return mx - mn
}

func main() {
	println(maxDiff(111))
}
