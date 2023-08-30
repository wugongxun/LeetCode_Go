package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v", minimumJumps([]int{128, 178, 147, 165, 63, 11, 150, 20, 158, 144, 136}, 61, 170, 135))
}

func minimumJumps(forbidden []int, a, b, x int) int {
	maxVal := 0
	forbiddenSet := make(map[int]bool)
	for _, f := range forbidden {
		//maxVal = max(maxVal, f)
		forbiddenSet[f] = true
	}
	//上限证明。。。(buhui)
	upper := max(maxVal+a, x) + b
	deque := [][3]int{{0, 1, 0}}
	visited := make(map[int]bool)
	visited[0] = true
	for len(deque) > 0 {
		pos, dir, step := deque[0][0], deque[0][1], deque[0][2]
		deque = deque[1:]
		if pos == x {
			return step
		}
		//前进
		nextPos, nextDir := pos+a, 1
		_, ok1 := visited[nextPos*nextDir]
		_, ok2 := forbiddenSet[nextPos]
		if nextPos >= 0 && nextPos <= upper && !ok1 && !ok2 {
			visited[nextPos*nextDir] = true
			deque = append(deque, [3]int{nextPos, nextDir, step + 1})
		}
		if dir == 1 {
			//向后
			nextPos, nextDir := pos-b, -1
			_, ok1 := visited[nextPos*nextDir]
			_, ok2 := forbiddenSet[nextPos]
			if nextPos >= 0 && nextPos <= upper && !ok1 && !ok2 {
				visited[nextPos*nextDir] = true
				deque = append(deque, [3]int{nextPos, nextDir, step + 1})
			}
		}
	}
	return -1
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
