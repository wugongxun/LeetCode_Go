package main

import "fmt"

func main() {
	lists := make([]*ListNode, 0)
	lists = append(lists, &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}})
	lists = append(lists, &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}})
	lists = append(lists, &ListNode{Val: 4, Next: &ListNode{Val: 6}})
	fmt.Printf("%v", mergeKLists(lists))
}

var n int

func mergeKLists(lists []*ListNode) *ListNode {
	n = len(lists)
	if n == 0 {
		return nil
	}
	return divide(lists, 0, n-1)
}

func divide(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	mid := (l + r) / 2
	return merge(divide(lists, l, mid), divide(lists, mid+1, r))
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	for l1 != nil {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	return res.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
