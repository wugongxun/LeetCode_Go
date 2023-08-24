package main

import (
	"sort"
)

func main() {

}

func reconstructQueue(people [][]int) (res [][]int) {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	for _, p := range people {
		idx := p[1]
		res = append(res[:idx], append([][]int{p}, res[idx:]...)...)
	}
	return res
}
