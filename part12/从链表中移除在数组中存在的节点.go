package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	res := &ListNode{-1, head}
	cur := res
	for cur.Next != nil {
		if set[cur.Next.Val] {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return res.Next
}

func main() {

}
