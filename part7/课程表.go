package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	//0表示未搜索，1表示搜索中，2表示已完成
	visited := make([]int, numCourses)
	valid := true
	adj := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adj[i] = make([]int, 0)
	}
	for _, prerequisite := range prerequisites {
		adj[prerequisite[1]] = append(adj[prerequisite[1]], prerequisite[0])
	}
	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = 1
		for _, next := range adj[i] {
			if visited[next] == 0 {
				dfs(next)
				if !valid {
					return
				}
			} else if visited[next] == 1 {
				valid = false
				return
			}
		}
		visited[i] = 2
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
