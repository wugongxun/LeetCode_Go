package main

import "fmt"

func lexicalOrder(n int) (res []int) {
	var dfs func(p int)
	dfs = func(p int) {
		i := 0
		if p == 0 {
			i = 1
		}
		for ; i <= 9; i++ {
			t := p*10 + i
			if t > n {
				break
			}
			res = append(res, t)
			dfs(t)
		}
	}
	dfs(0)
	return
}

func main() {
	fmt.Printf("%v", lexicalOrder(100))
}
