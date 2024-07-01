package main

func main() {

}

func maximalPathQuality(values []int, edges [][]int, maxTime int) (res int) {
	n := len(values)
	type edge struct{ next, time int }
	adj := make([][]edge, n)
	for _, e := range edges {
		u, v, t := e[0], e[1], e[2]
		adj[u] = append(adj[u], edge{v, t})
		adj[v] = append(adj[v], edge{u, t})
	}
	visited := make([]bool, n)
	visited[0] = true
	var dfs func(cur, sumTime, sumValue int)
	dfs = func(cur, sumTime, sumValue int) {
		if cur == 0 {
			res = max(res, sumValue)
		}
		for _, e := range adj[cur] {
			next, nextTime := e.next, e.time
			if nextTime+sumTime > maxTime {
				continue
			}
			if visited[next] {
				dfs(next, sumTime+nextTime, sumValue)
			} else {
				visited[next] = true
				dfs(next, sumTime+nextTime, sumValue+values[next])
				visited[next] = false
			}
		}
	}
	dfs(0, 0, values[0])
	return res
}
