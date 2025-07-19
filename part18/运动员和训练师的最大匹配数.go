package main

import "slices"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	slices.Sort(players)
	slices.Sort(trainers)
	j := 0
	for i, p := range players {
		for j < len(trainers) && trainers[j] < p {
			j++
		}
		if j == len(trainers) {
			return i
		}
		j++
	}
	return len(players)
}

func main() {
	println(matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8}))
}
