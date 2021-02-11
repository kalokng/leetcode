package lib

import (
	"fmt"
	"strings"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func (n *Node) String() string {
	if n == nil {
		return "(nil)"
	}
	var s strings.Builder
	for n != nil {
		fmt.Fprint(&s, "->", n.Val)
		rand := "(nil)"
		if n.Random != nil {
			rand = fmt.Sprint("(", n.Random.Val, ")")
		}
		s.WriteString(rand)
		n = n.Next
	}
	return s.String()[2:]
}

func NewNode(v ...interface{}) *Node {
	if len(v) == 0 {
		return nil
	}
	if len(v)%2 != 0 {
		panic("argument not in pairs")
	}
	last := &Node{}
	ns := make([]*Node, len(v)/2)
	for i := 0; i < len(v); i += 2 {
		node := &Node{
			Val: v[i].(int),
		}
		last.Next = node
		ns[i/2] = node
		last = node
	}
	for i := 1; i < len(v); i += 2 {
		var rand *Node
		if j, ok := v[i].(int); ok {
			rand = ns[j]
		}
		ns[i/2].Random = rand
	}
	return ns[0]
}
