package main

func closestMeetingNode(edges []int, x, y int) int {
	n := len(edges)
	res := n
	visX, visY := make([]bool, n), make([]bool, n)
	for !visX[x] || !visY[y] {
		visX[x], visY[y] = true, true
		if visY[x] {
			res = x
		}
		if visX[y] {
			res = min(res, y)
		}
		if res < n {
			return res
		}
		if edges[x] >= 0 {
			x = edges[x]
		}
		if edges[y] >= 0 {
			y = edges[y]
		}
	}
	return -1
}

//func closestMeetingNode(edges []int, node1 int, node2 int) int {
//	n := len(edges)
//	minDis, res := n, -1
//
//	calcDis := func(x int) []int {
//		dis := make([]int, n)
//		for i := range dis {
//			dis[i] = n
//		}
//		for d := 0; x != -1 && dis[x] == n; x = edges[x] {
//			dis[x] = d
//			d++
//		}
//		return dis
//	}
//
//	d1 := calcDis(node1)
//	d2 := calcDis(node2)
//
//	for i := 0; i < n; i++ {
//		d := max(d1[i], d2[i])
//		if d < minDis {
//			minDis, res = d, i
//		}
//	}
//	return res
//}
