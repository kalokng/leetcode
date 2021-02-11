package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	fmt.Println(findLHS([]int{1, 2, 3, 4}))
	fmt.Println(findLHS([]int{1, 1, 1, 1}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLHS(nums []int) int {
	mx := 0
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if b, ok := m[k+1]; ok {
			mx = max(mx, v+b)
		}
	}
	return mx
}
