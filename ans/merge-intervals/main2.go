package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 3}, {4, 10}, {3, 4}, {15, 18}}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	i := 0
	for j := 1; j < len(intervals); j++ {
		p := intervals[j]
		pp := intervals[i]
		if pp[1] >= p[0] {
			pp[1] = max(pp[1], p[1])
			continue
		}
		i++
		intervals[i] = p
	}
	return intervals[:i+1]
}
