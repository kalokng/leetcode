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
	p := tr(root, 0)
	var r [][]int
	for i, v := range p.v {
		if len(v) == 0 {
			continue
		}
		p.i = i
		sort.Sort(&p)
		r = append(r, p.v[i])
	}
	return r
}

type pair struct {
	v [][]int
	y [][]int
	s int
	i int
}

func (p *pair) Len() int { return len(p.v[p.i]) }
func (p *pair) Less(i, j int) bool {
	if p.y[p.i][i] == p.y[p.i][j] {
		return p.v[p.i][i] < p.v[p.i][j]
	}
	return p.y[p.i][i] < p.y[p.i][j]
}
func (p *pair) Swap(i, j int) {
	p.v[p.i][i], p.v[p.i][j] = p.v[p.i][j], p.v[p.i][i]
	p.y[p.i][i], p.y[p.i][j] = p.y[p.i][j], p.y[p.i][i]
}

func tr(root *TreeNode, lvl int) pair {
	if root == nil {
		return pair{
			v: nil,
			y: nil,
			s: 0,
		}
	}
	l := tr(root.Left, lvl+1)
	r := tr(root.Right, lvl+1)
	l.s--
	r.s++
	min := l.s
	var v, y [][]int
	if r.s < min {
		min = r.s
	}
	max := r.s + len(r.v)
	if l.s+len(l.v) > max {
		max = l.s + len(l.v)
	}
	for i := min; i < max; i++ {
		var vv, yy []int
		if i >= l.s && i-l.s < len(l.v) {
			vv = append(vv, l.v[i-l.s]...)
			yy = append(yy, l.y[i-l.s]...)
		}
		if i >= r.s && i-r.s < len(r.v) {
			vv = append(vv, r.v[i-r.s]...)
			yy = append(yy, r.y[i-r.s]...)
		}
		if i == 0 {
			vv = append(vv, root.Val)
			yy = append(yy, lvl)
		}
		v = append(v, vv)
		y = append(y, yy)
	}
	return pair{
		v: v,
		y: y,
		s: min,
	}
}
