package lib

import (
	"fmt"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
	loop bool
}

func (n *ListNode) String() string {
	if n == nil {
		return "(nil)"
	}
	var s strings.Builder
	for n != nil && !n.loop {
		fmt.Fprint(&s, "->", n.Val)
		n = n.Next
	}
	if n != nil {
		fmt.Fprint(&s, "->(", n.Val)
		n = n.Next
		for n != nil && !n.loop {
			fmt.Fprint(&s, "->", n.Val)
			n = n.Next
		}
		s.WriteString("->loop)")
	}
	return s.String()[2:]
}

func NewList(n ...int) *ListNode {
	var root *ListNode
	for i := len(n) - 1; i >= 0; i-- {
		root = &ListNode{
			Val:  n[i],
			Next: root,
		}
	}
	return root
}

func NewCycleList(pos int, n ...int) *ListNode {
	if pos < 0 || pos >= len(n) {
		return NewList(n...)
	}
	var root = &ListNode{loop: true}
	tail := root
	for i := len(n) - 1; i >= 0; i-- {
		if pos == i {
			tail.Val = n[i]
			tail.Next = root
			root = tail
			continue
		}
		root = &ListNode{
			Val:  n[i],
			Next: root,
		}
	}
	return root
}
