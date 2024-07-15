package main

import "slices"

func accountsMerge(accounts [][]string) (res [][]string) {
	n := len(accounts)
	emailToIdx := make(map[string][]int)
	for i, account := range accounts {
		for _, email := range account[1:] {
			emailToIdx[email] = append(emailToIdx[email], i)
		}
	}
	vis := make([]bool, n)
	emailSet := make(map[string]bool)
	var dfs func(i int)
	dfs = func(i int) {
		vis[i] = true
		for _, email := range accounts[i][1:] {
			if emailSet[email] {
				continue
			}
			emailSet[email] = true
			for _, next := range emailToIdx[email] {
				if !vis[next] {
					dfs(next)
				}
			}
		}
	}
	for i, b := range vis {
		if b {
			continue
		}
		clear(emailSet)
		dfs(i)
		list := make([]string, 1, len(emailSet)+1)
		list[0] = accounts[i][0]
		for email := range emailSet {
			list = append(list, email)
		}
		slices.Sort(list[1:])
		res = append(res, list)
	}
	return
}
