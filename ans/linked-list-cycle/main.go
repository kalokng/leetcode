package main

import (
	"fmt"

	. "github.com/kalokng/leetcode/lib"
)

func main() {
	fmt.Println(hasCycle(NewCycleList(1, 3, 2, 0, -4)))
	fmt.Println(hasCycle(NewCycleList(0, 1, 2)))
	fmt.Println(hasCycle(NewList(1)))
	fmt.Println(hasCycle(NewList()))
	fmt.Println(hasCycle(NewCycleList(0, 1)))
	//fmt.Println(hasCycle(
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	h := head.Next
	for h != nil {
		if h.Next == nil {
			return false
		}
		head = head.Next
		h = h.Next.Next
		if head == h {
			return true
		}
	}
	return h != nil
}
