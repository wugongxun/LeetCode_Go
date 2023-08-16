package main

import (
	"fmt"
)

func main() {
	root1 := &TreeNode{1, &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, nil}
	root2 := &TreeNode{1, nil, &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	fmt.Printf("%v", mergeTrees(root1, root2))
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
