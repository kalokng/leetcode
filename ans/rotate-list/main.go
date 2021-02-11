package main

import (
	"fmt"

	. "github.com/kalokng/leetcode/lib"
)

func main() {
	fmt.Println(rotateRight(NewList(1, 2, 3, 4, 5), 1))
	fmt.Println(rotateRight(NewList(1, 2, 3, 4, 5), 2))
	fmt.Println(rotateRight(NewList(1, 2, 3, 4, 5), 3))
	fmt.Println(rotateRight(NewList(1, 2, 3, 4, 5), 4))
	fmt.Println(rotateRight(NewList(1, 2, 3, 4, 5), 5))
	fmt.Println(rotateRight(NewList(1, 2, 3, 4, 5), 6))
	fmt.Println(rotateRight(NewList(1, 2, 3, 4, 5), 7))
	fmt.Println(rotateRight(NewList(1, 2, 3, 4, 5), 8))
	fmt.Println(rotateRight(NewList(1, 2, 3, 4, 5), 9))
	fmt.Println(rotateRight(NewList(1, 2, 3, 4, 5), 10))
	fmt.Println(rotateRight(NewList(1), 1))
	fmt.Println(rotateRight(NewList(1, 2), 1))
	fmt.Println(rotateRight(NewList(1, 2), 2))
	fmt.Println(rotateRight(NewList(), 1))
	fmt.Println(rotateRight(NewList(1), 100))
	fmt.Println(rotateRight(NewList(), 100))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := head
	n := 1
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	k = n - k%n
	cur := head
	for i := 1; i < k; i++ {
		cur = cur.Next
	}
	tail.Next = head
	head = cur.Next
	cur.Next = nil
	return head
}
