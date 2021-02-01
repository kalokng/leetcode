package main

import (
	"fmt"
	"sort"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	m := make(map[int]int)
	for _, l := range lists {
		for l != nil {
			m[l.Val]++
			l = l.Next
		}
	}
	var l = make([]int, 0, len(m))
	for k := range m {
		l = append(l, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(l)))
	fmt.Println(l)

	var node *ListNode
	for _, i := range l {
		for j := 0; j < m[i]; j++ {
			node = &ListNode{
				Val:  i,
				Next: node,
			}
		}
	}
	return node
}

func main() {
	l1 := newList([]int{1, 3, 4, 6, 8, 9, 12})
	l2 := newList([]int{1, 2, 5, 7, 11, 21, 24})
	l3 := newList([]int{-4, 0, 4, 7, 10, 14, 22, 29})
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(mergeKLists([]*ListNode{l1, l2, l3}))
	fmt.Println(mergeKLists(nil))
	fmt.Println(mergeKLists([]*ListNode{nil, nil, nil}))
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
