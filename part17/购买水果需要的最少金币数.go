package main

func minimumCoins(prices []int) int {
	n := len(prices)
	type pair struct{ i, f int }
	q := []pair{{n + 1, 0}}
	for i := n; i > 0; i-- {
		for q[0].i > i*2+1 {
			q = q[1:]
		}
		f := q[0].f + prices[i-1]
		for f <= q[len(q)-1].f {
			q = q[:len(q)-1]
		}
		q = append(q, pair{i, f})
	}
	return q[len(q)-1].f
	//n := len(prices)
	//for i := (n+1)/2 - 1; i > 0; i-- {
	//	prices[i-1] += slices.Min(prices[i : i*2+1])
	//}
	//return prices[0]
}

func main() {
	println(minimumCoins([]int{3, 1, 2}))
}
