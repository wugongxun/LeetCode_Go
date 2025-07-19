package main

import (
	"fmt"
	"slices"
	"strings"
)

func removeSubfolders(folder []string) []string {
	slices.Sort(folder)
	res := folder[:1]
	for _, f := range folder[1:] {
		last := res[len(res)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			res = append(res, f)
		}
	}
	return res
}

func main() {
	fmt.Printf("%v", removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
}
