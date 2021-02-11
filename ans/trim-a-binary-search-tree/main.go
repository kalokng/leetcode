package main

import (
	"fmt"

	. "github.com/kalokng/leetcode/lib"
)

func main() {
	fmt.Println(trimBST(NewTree(1, 0, 2), 1, 2))
	fmt.Println(trimBST(NewTree(3, 0, 4, nil, 2, nil, nil, 1), 1, 3))
	fmt.Println(trimBST(NewTree(1), 1, 2))
	fmt.Println(trimBST(NewTree(1, nil, 2), 1, 3))
	fmt.Println(trimBST(NewTree(1, nil, 2), 2, 4))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	root.Left = trimLow(root.Left, low)
	root.Right = trimHigh(root.Right, high)
	return root
}
func trimLow(root *TreeNode, low int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimLow(root.Right, low)
	}
	root.Left = trimLow(root.Left, low)
	return root
}
func trimHigh(root *TreeNode, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > high {
		return trimHigh(root.Left, high)
	}
	root.Right = trimHigh(root.Right, high)
	return root
}
