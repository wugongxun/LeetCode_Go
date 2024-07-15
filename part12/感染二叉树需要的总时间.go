package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	fa := make(map[*TreeNode]*TreeNode)
	var startNode *TreeNode
	var dfs func(*TreeNode, *TreeNode)
	dfs = func(node, form *TreeNode) {
		if node == nil {
			return
		}
		fa[node] = form
		if node.Val == start {
			startNode = node
		}
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)
	var maxDepth func(*TreeNode, *TreeNode) int
	maxDepth = func(node, form *TreeNode) int {
		if node == nil {
			return -1
		}
		res := -1
		if node.Left != form {
			res = max(res, maxDepth(node.Left, node))
		}
		if node.Right != form {
			res = max(res, maxDepth(node.Right, node))
		}
		if fa[node] != form {
			res = max(res, maxDepth(fa[node], node))
		}
		return res + 1
	}
	return maxDepth(startNode, startNode)
}

func main() {
	print(amountOfTime(&TreeNode{1, &TreeNode{5, nil, &TreeNode{4, &TreeNode{Val: 9}, &TreeNode{Val: 2}}}, &TreeNode{3, &TreeNode{Val: 10}, &TreeNode{Val: 6}}}, 3))
}
