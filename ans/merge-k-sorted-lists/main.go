package main

import (
	"fmt"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var c = make(chan int)
	close(c)
	for _, l := range lists {
		c = merge2(c, chanList(l))
	}

	var root *ListNode
	var cur = root
	for v := range c {
		last := cur
		cur = &ListNode{
			Val: v,
		}
		if last == nil {
			root = cur
			continue
		}
		last.Next = cur
	}
	return root
}

func merge2(a, b chan int) chan int {
	c := make(chan int, 1)
	go func() {
		defer close(c)
		av, aok := <-a
		bv, bok := <-b
		for aok && bok {
			if av > bv {
				c <- bv
				bv, bok = <-b
			} else {
				c <- av
				av, aok = <-a
			}
		}
		if aok {
			c <- av
		}
		if bok {
			c <- bv
		}
		for v := range a {
			c <- v
		}
		for v := range b {
			c <- v
		}
	}()
	return c
}

func chanList(n *ListNode) chan int {
	c := make(chan int, 1)
	go func() {
		defer close(c)
		for n != nil {
			c <- n.Val
			n = n.Next
		}
	}()
	return c
}

func main() {
	l1 := NewList([]int{1, 3, 5, 7, 9})
	l2 := NewList([]int{2, 4, 6, 8})
	l3 := NewList([]int{1, 1, 2, 3, 5, 8, 13})
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
