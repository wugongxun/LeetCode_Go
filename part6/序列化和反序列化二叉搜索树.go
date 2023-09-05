package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	codec := Constructor()
	root := TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}
	serialize := codec.serialize(&root)
	fmt.Printf("%v", serialize)
	fmt.Printf("%v", codec.deserialize(serialize))
}

type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (Codec) serialize(root *TreeNode) string {
	arr := make([]string, 0)
	var postOrder func(*TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		arr = append(arr, strconv.Itoa(node.Val))
	}
	postOrder(root)
	return strings.Join(arr, ",")
}

// Deserializes your encoded data to tree.
func (Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, ",")
	var construct func(int, int) *TreeNode
	construct = func(lower int, upper int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(arr[len(arr)-1])
		if val < lower || val > upper {
			return nil
		}
		arr = arr[:len(arr)-1]
		return &TreeNode{Val: val, Right: construct(val, upper), Left: construct(lower, val)}
	}
	return construct(math.MinInt32, math.MaxInt32)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
