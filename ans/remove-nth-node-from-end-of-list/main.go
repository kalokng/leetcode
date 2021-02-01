package main

import (
	"fmt"

	. "github.com/kalokng/leetcode/lib"
)

func main() {
	fmt.Println(removeNthFromEnd(NewList(1, 2, 3, 4, 5), 1))
	fmt.Println(removeNthFromEnd(NewList(1, 2, 3, 4, 5), 2))
	fmt.Println(removeNthFromEnd(NewList(1, 2, 3, 4, 5), 3))
	fmt.Println(removeNthFromEnd(NewList(1, 2, 3, 4, 5), 4))
	fmt.Println(removeNthFromEnd(NewList(1, 2, 3, 4, 5), 5))
	fmt.Println(removeNthFromEnd(NewList(1), 1))
	fmt.Println(removeNthFromEnd(NewList(1, 2), 1))
	fmt.Println(removeNthFromEnd(NewList(1, 2), 2))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root := &ListNode{Next: head}
	tail := head
	for i := 0; i < n; i++ {
		tail = tail.Next
	}
	cur := root
	for tail != nil {
		tail = tail.Next
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return root.Next
}
