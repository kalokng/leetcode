package main

import "fmt"

func main() {
	fmt.Println(getCoprimes([]int{2, 3, 3, 2}, [][]int{{0, 1}, {1, 2}, {1, 3}}))
	fmt.Println(getCoprimes([]int{5, 6, 10, 2, 3, 6, 15}, [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}))
}

func gcd(x, y int) int {
	if x > y {
		x, y = y, x
	}
	for {
		if x == 0 {
			return y
		}
		x, y = y%x, x
	}
}

type node struct {
	idx  int
	val  int
	lv   int
	p    int
	edge []int
}

func getCoprimes(nums []int, edges [][]int) []int {
	gm := make([][]int, 51)
	for i := 1; i <= 50; i++ {
		for j := 1; j <= 50; j++ {
			if g := gcd(i, j); g == 1 {
				gm[i] = append(gm[i], j)
			}
		}
	}

	tree := make([]*node, len(nums))
	for i, v := range nums {
		tree[i] = &node{
			idx: i,
			val: v,
		}
	}
	for _, e := range edges {
		tree[e[0]].edge = append(tree[e[0]].edge, e[1])
		tree[e[1]].edge = append(tree[e[1]].edge, e[0])
	}
	l := []int{0}
	for len(l) > 0 {
		var nl []int
		for _, i := range l {
			n := tree[i]
			p := n.p
			for i := 0; i < len(n.edge); i++ {
				if n.edge[i] == p {
					n.edge = append(n.edge[:i], n.edge[i+1:]...)
					break
				}
			}
			for _, e := range n.edge {
				tree[e].p = i
				tree[e].lv = n.lv + 1
			}
			nl = append(nl, n.edge...)
		}
		l = nl
	}
	tree[0].p = -1
	//for i, v := range tree {
	//	fmt.Println(i, v)
	//}

	lv := make(map[int]int)
	var dfs func(e *node)
	dfs = func(e *node) {
		g := gm[e.val]
		mx := -1
		c := -1
		for _, i := range g {
			if a, ok := lv[i]; ok {
				if tree[a].lv > mx {
					mx = tree[a].lv
					c = a
				}
			}
		}
		nums[e.idx] = c
		o, ok := lv[e.val]
		lv[e.val] = e.idx
		for _, x := range e.edge {
			dfs(tree[x])
		}
		if ok {
			lv[e.val] = o
		} else {
			delete(lv, e.val)
		}
	}

	dfs(tree[0])
	return nums
}
