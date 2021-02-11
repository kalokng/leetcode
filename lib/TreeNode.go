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
	rv, ok := node[0].(int)
	if !ok {
		return nil
	}
	root := &TreeNode{
		Val:   rv,
		Left:  nil,
		Right: nil,
	}

	var q = []*TreeNode{root}

	l := true
	for _, v := range node[1:] {
		var cur *TreeNode
		switch d := v.(type) {
		case int:
			cur = &TreeNode{
				Val:   d,
				Left:  nil,
				Right: nil,
			}
			q = append(q, cur)
		default:
		}
		if l {
			q[0].Left = cur
			l = false
		} else {
			q[0].Right = cur
			q = q[1:]
			l = true
		}
	}
	return root
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
