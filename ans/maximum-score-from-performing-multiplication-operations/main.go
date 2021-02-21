package main

import "fmt"

func main() {
	fmt.Println(maximumScore([]int{1, 2, 3}, []int{3, 2, 1}))
	fmt.Println(maximumScore([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maximumScore(nums []int, multipliers []int) int {
	c := make(map[P]int)
	return scr(c, nums, multipliers, 0, 0, len(nums)-1)
}

type P struct {
	i, j int
}

func scr(c map[P]int, n []int, m []int, mi, i, j int) (res int) {
	if v, ok := c[P{i, j}]; ok {
		return v
	}
	defer func() {
		c[P{i, j}] = res
	}()
	if i == j {
		return n[i] * m[mi]
	}
	var o1, o2 int
	if mi < len(m)-1 {
		o1 = scr(c, n, m, mi+1, i+1, j)
		o2 = scr(c, n, m, mi+1, i, j-1)
	}
	o1 += n[i] * m[mi]
	o2 += n[j] * m[mi]
	return max(o1, o2)
}
