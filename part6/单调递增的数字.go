package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v", monotoneIncreasingDigits(12342134))
}

func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			j := i - 1
			for j > 0 && s[j] <= s[j-1] {
				j--
			}
			s[j]--
			for j++; j < len(s); j++ {
				s[j] = '9'
			}
			break
		}
	}
	res, _ := strconv.Atoi(string(s))
	return res
}
