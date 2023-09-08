package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v", repairCars([]int{4, 2, 3, 1}, 10))
}

func repairCars(ranks []int, cars int) int64 {
	lower, upper := 0, ranks[0]*cars*cars

	var calculate = func(time int) int {
		res := 0
		for _, rank := range ranks {
			res += int(math.Sqrt(float64(time / rank)))
		}
		return res
	}

	for lower < upper {
		mid := (lower + upper) >> 1
		if calculate(mid) >= cars {
			upper = mid
		} else {
			lower = mid + 1
		}
	}
	return int64(lower)
}
