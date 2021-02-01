package main

import (
	"container/heap"
	"fmt"

	. "github.com/kalokng/leetcode/lib"
)

type ListHeap []*ListNode

func (h ListHeap) Len() int           { return len(h) }
func (h ListHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	idx := make(ListHeap, 0, len(lists))
	for _, l := range lists {
		if l == nil {
			continue
		}
		idx = append(idx, l)
	}
	heap.Init(&idx)
	root := &ListNode{Next: nil}
	cur := root
	for len(idx) > 0 {
		node := idx[0]
		cur.Next = node
		cur = node
		if node.Next == nil {
			heap.Pop(&idx)
			continue
		}
		idx[0] = node.Next
		heap.Fix(&idx, 0)
	}
	return root.Next
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
	fmt.Println(mergeKLists([]*ListNode{nil, nil, nil, l1}))
}
