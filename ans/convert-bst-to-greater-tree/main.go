package main

import (
	"fmt"

	. "github.com/kalokng/leetcode/lib"
)

func main() {
	t := NewTree(4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8)
	fmt.Println(convertBST(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	weight(root, 0)
	return root
}

func weight(root *TreeNode, w int) int {
	if root.Right != nil {
		w = weight(root.Right, w)
	}
	w += root.Val
	root.Val = w
	if root.Left != nil {
		w = weight(root.Left, w)
	}
	return w
}
