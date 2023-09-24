package main

func main() {
	tree := Constructor([]int{-1, 0, 0, 1, 1, 2, 2})
	tree.Lock(2, 2)
	tree.Unlock(2, 3)
	tree.Unlock(2, 2)
	tree.Lock(4, 5)
	tree.Upgrade(0, 1)
	tree.Lock(0, 1)
}

type LockingTree struct {
	parent   []int
	lock     []int
	children [][]int
}

func Constructor(parent []int) (lockingTree LockingTree) {
	lockingTree.parent = parent
	lockingTree.lock = make([]int, len(parent))
	children := make([][]int, len(parent))
	for i, p := range parent {
		if p != -1 {
			children[p] = append(children[p], i)
		}
	}
	lockingTree.children = children
	return
}

func (lockingTree *LockingTree) Lock(num int, user int) bool {
	if lockingTree.lock[num] == 0 {
		lockingTree.lock[num] = user
		return true
	}
	return false
}

func (lockingTree *LockingTree) Unlock(num int, user int) bool {
	if lockingTree.lock[num] == user {
		lockingTree.lock[num] = 0
		return true
	}
	return false
}

func (lockingTree *LockingTree) Upgrade(num int, user int) bool {
	if lockingTree.lock[num] != 0 {
		return false
	}
	//判断祖先是否有上锁
	res := false
	cur := lockingTree.parent[num]
	for cur != -1 {
		if lockingTree.lock[cur] != 0 {
			res = true
			break
		}
		cur = lockingTree.parent[cur]
	}
	//判断子孙是否有上锁，并且解锁
	var dfs func(num int) bool
	dfs = func(num int) bool {
		res := lockingTree.lock[num] != 0
		lockingTree.lock[num] = 0
		for _, child := range lockingTree.children[num] {
			if dfs(child) {
				res = true
			}
		}
		return res
	}
	if !res && dfs(num) {
		lockingTree.lock[num] = user
		return true
	}
	return false
}
