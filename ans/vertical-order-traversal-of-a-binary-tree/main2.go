package main

import (
	"fmt"
	"sort"

	. "github.com/kalokng/leetcode/lib"
)

func main() {
	fmt.Println(verticalTraversal(NewTree(3, 9, 20, nil, nil, 15, 7)))
	fmt.Println(verticalTraversal(NewTree(1, 2, 3, 4, 5, 6, 7)))
	fmt.Println(verticalTraversal(NewTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)))
	fmt.Println(verticalTraversal(NewTree(7, 6, 5, 4, 3, 2, 1)))
	fmt.Println(verticalTraversal(NewTree(15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)))
	fmt.Println(verticalTraversal(NewTree(1, nil, 2, nil, nil, 3, 4)))
	fmt.Println(verticalTraversal(NewTree(1)))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {
	m := make(map[pair][]int)
	tr(m, root, 0, 0)
	var p = make([]pair, 0, len(m))
	for k := range m {
		p = append(p, k)
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			return p[i].y < p[j].y
		}
		return p[i].x < p[j].x
	})
	x := p[0].x
	var v []int
	var vv [][]int
	for _, k := range p {
		if k.x != x {
			vv = append(vv, v)
			x = k.x
			v = nil
		}
		y := m[k]
		sort.Ints(y)
		v = append(v, y...)
	}
	vv = append(vv, v)
	return vv
}

type pair struct {
	x, y int
}

func tr(m map[pair][]int, root *TreeNode, x, y int) {
	p := pair{x, y}
	m[p] = append(m[p], root.Val)
	if root.Left != nil {
		tr(m, root.Left, x-1, y+1)
	}
	if root.Right != nil {
		tr(m, root.Right, x+1, y+1)
	}
}
