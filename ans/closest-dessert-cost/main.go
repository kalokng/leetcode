package main

import "fmt"

func main() {
	fmt.Println(closestCost([]int{1, 7}, []int{3, 4}, 10))
	fmt.Println(closestCost([]int{2, 3}, []int{4, 5, 100}, 18))
	fmt.Println(closestCost([]int{3, 10}, []int{2, 5}, 9))
	fmt.Println(closestCost([]int{10}, []int{1}, 1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	m := map[int][]int{}
	md, mn := cc(m, toppingCosts, 0, target, baseCosts[0])
	for i := 1; i < len(baseCosts); i++ {
		m := map[int][]int{}
		d, v := cc(m, toppingCosts, 0, target, baseCosts[i])
		if d < md || (d == md && v < mn) {
			md, mn = d, v
		}
	}
	return mn
}

func cc(m map[int][]int, t []int, i, tgt int, s int) (d int, v int) {
	if len(t) == i {
		return abs(tgt - s), s
	}
	d0, v0 := cc(m, t, i+1, tgt, s)
	d1, v1 := cc(m, t, i+1, tgt, s+t[i])
	if d1 < d0 || (d1 == d0 && v1 < v0) {
		d0, v0 = d1, v1
	}
	d2, v2 := cc(m, t, i+1, tgt, s+t[i]*2)
	if d2 < d0 || (d2 == d0 && v2 < v0) {
		return d2, v2
	}
	return d0, v0
}
