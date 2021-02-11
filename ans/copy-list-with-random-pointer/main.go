package main

import (
	"fmt"

	. "github.com/kalokng/leetcode/lib"
)

func main() {
	n := NewNode(7, nil, 13, 0, 11, 4, 10, 2, 1, 0)
	fmt.Println(n)
	fmt.Println(copyRandomList(n))
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	root := &Node{}
	last := root
	for head != nil {
		n := &Node{
			Val:    head.Val,
			Random: head.Random,
		}
		last.Next = n
		last = n
		m[head] = n
		head = head.Next
	}
	last = root.Next
	for last != nil {
		if last.Random != nil {
			last.Random = m[last.Random]
		}
		last = last.Next
	}
	return root.Next
}
