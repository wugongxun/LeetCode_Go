package main

func longestPalindrome(words []string) int {
	cnt := [26][26]int{}
	for _, w := range words {
		cnt[w[0]-'a'][w[1]-'a']++
	}

	res, odd := 0, 0
	for i := range cnt {
		c := cnt[i][i]
		res += c - c%2
		odd |= c % 2
		for j := i + 1; j < 26; j++ {
			res += min(cnt[i][j], cnt[j][i]) * 2
		}
	}
	return (res + odd) * 2

	//cnt := make(map[string]int)
	//for _, word := range words {
	//	cnt[word]++
	//}
	//res, odd := 0, 0
	//for w, c := range cnt {
	//	if w[0] == w[1] {
	//		res += c - c%2
	//		odd |= c % 2
	//	} else if w[0] > w[1] {
	//		res += min(c, cnt[string([]byte{w[1], w[0]})]) * 2
	//	}
	//}
	//return (res + odd) * 2
}

func main() {
	println(longestPalindrome([]string{"lc", "cl", "gg"}))
}
