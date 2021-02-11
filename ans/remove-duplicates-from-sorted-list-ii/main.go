package main

import (
	"fmt"

	. "github.com/kalokng/leetcode/lib"
)

func main() {
	fmt.Println(deleteDuplicates(NewList(1, 2, 3, 3, 4, 4, 5)))
	fmt.Println(deleteDuplicates(NewList(1, 2, 3, 3, 3, 4, 4, 5)))
	fmt.Println(deleteDuplicates(NewList(1, 1, 1, 2, 3)))
	fmt.Println(deleteDuplicates(nil))
	fmt.Println(deleteDuplicates(NewList(1, 1, 1)))
	fmt.Println(deleteDuplicates(NewList(1, 1, 2, 2, 3, 3)))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	root := &ListNode{Next: head}
	last := root
	cur := head
	for cur.Next != nil {
		if cur.Val != cur.Next.Val {
			last = cur
			cur = cur.Next
			continue
		}
		v := cur.Val
		cur = cur.Next
		for cur != nil && cur.Val == v {
			cur = cur.Next
		}
		last.Next = cur
		if cur == nil {
			break
		}
	}
	return root.Next
}
