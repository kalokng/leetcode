package lib

import (
	"fmt"
	"strings"
)

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

func NewList(n []int) *ListNode {
	var root *ListNode
	for i := len(n) - 1; i >= 0; i-- {
		root = &ListNode{
			Val:  n[i],
			Next: root,
		}
	}
	return root
}
