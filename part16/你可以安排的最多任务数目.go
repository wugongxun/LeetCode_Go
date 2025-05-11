package main

import (
	"slices"
	"sort"
)

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	slices.Sort(tasks)
	slices.Sort(workers)
	n, m := len(tasks), len(workers)
	return sort.Search(min(n, m), func(k int) bool {
		k++
		i, p := 0, pills
		var validTasks []int
		for _, w := range workers[m-k:] {
			for ; i < k && tasks[i] <= w+strength; i++ {
				validTasks = append(validTasks, tasks[i])
			}
			if len(validTasks) == 0 {
				return true
			}
			if w >= validTasks[0] {
				validTasks = validTasks[1:]
				continue
			}
			if p == 0 {
				return true
			}
			p--
			validTasks = validTasks[:len(validTasks)-1]
		}
		return false
	})
}

func main() {
	println(maxTaskAssign([]int{3, 2, 1}, []int{0, 3, 3}, 1, 1))
}
