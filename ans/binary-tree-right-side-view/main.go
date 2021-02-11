package main

import (
	"fmt"

	. "github.com/kalokng/leetcode/lib"
)

func main() {
	fmt.Println(rightSideView(NewTree(1, 2, 3, nil, 5, 7, 4, 6)))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	return v(root, 0)
}

func v(root *TreeNode, skip int) []int {
	if root == nil {
		return nil
	}
	var res []int
	if skip == 0 {
		res = append(res, root.Val)
	} else {
		skip--
	}
	r := v(root.Right, skip)
	l := v(root.Left, skip+len(r))
	res = append(res, r...)
	return append(res, l...)
}
