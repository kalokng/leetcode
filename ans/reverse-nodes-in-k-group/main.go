package main

import (
	"fmt"
	"strings"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	supRoot := &ListNode{Next: head}
	last := supRoot
	cur := head
outloop:
	for cur != nil {
		var i int
		for i = 0; i < k && cur != nil; i++ {
			cur = cur.Next
		}
		if i < k {
			break outloop
		}
		next := cur
		tail := next
		cur = last.Next
		for cur != next {
			cur, cur.Next, tail = cur.Next, tail, cur
		}
		cur, last, last.Next = next, last.Next, tail
	}
	return supRoot.Next
}

func main() {
	fmt.Println(reverseKGroup(newList([]int{1, 2, 3, 4, 5}), 0))
	fmt.Println(reverseKGroup(newList([]int{1, 2, 3, 4, 5}), 1))
	fmt.Println(reverseKGroup(newList([]int{1, 2, 3, 4, 5}), 2))
	fmt.Println(reverseKGroup(newList([]int{1, 2, 3, 4, 5}), 3))
	fmt.Println(reverseKGroup(newList([]int{1, 2, 3, 4, 5}), 4))
	fmt.Println(reverseKGroup(newList([]int{1, 2, 3, 4, 5}), 5))
	fmt.Println(reverseKGroup(newList([]int{1, 2, 3, 4, 5}), 6))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return "(nil)"
	}
	var s strings.Builder
	s.WriteString(fmt.Sprint(n.Val))
	n = n.Next
	for n != nil {
		s.WriteString(fmt.Sprint("->", n.Val))
		n = n.Next
	}
	return s.String()
}

func newList(n []int) *ListNode {
	var root *ListNode
	for i := len(n) - 1; i >= 0; i-- {
		root = &ListNode{
			Val:  n[i],
			Next: root,
		}
	}
	return root
}
