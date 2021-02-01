package main

import (
	"fmt"
	"sort"

	. "github.com/kalokng/leetcode/lib"
)

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
	l1 := NewList(1, 3, 4, 6, 8, 9, 12)
	l2 := NewList(1, 2, 5, 7, 11, 21, 24)
	l3 := NewList(-4, 0, 4, 7, 10, 14, 22, 29)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(mergeKLists([]*ListNode{l1, l2, l3}))
	fmt.Println(mergeKLists(nil))
	fmt.Println(mergeKLists([]*ListNode{nil, nil, nil}))
}
