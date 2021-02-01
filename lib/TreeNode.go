package lib

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(node ...interface{}) *TreeNode {
	if len(node) == 0 {
		return nil
	}
	var l = make([]*TreeNode, len(node))
	for i, v := range node {
		var cur *TreeNode
		switch d := v.(type) {
		case int:
			cur = &TreeNode{
				Val:   d,
				Left:  nil,
				Right: nil,
			}
		default:
		}
		if i > 0 {
			p := (i - 1) / 2
			if l[p] != nil {
				left := i%2 == 1
				if left {
					l[p].Left = cur
				} else {
					l[p].Right = cur
				}
			}
		}
		l[i] = cur
	}
	return l[0]
}

func (tree *TreeNode) String() string {
	if tree == nil {
		return "nil"
	}
	if tree.Left == nil && tree.Right == nil {
		return fmt.Sprintf("(%d)", tree.Val)
	}
	return fmt.Sprintf("(%d: %s %s)", tree.Val, tree.Left.String(), tree.Right.String())
}
