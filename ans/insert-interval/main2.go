package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	fmt.Println(insert([][]int{}, []int{5, 7}))
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 3}))
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 7}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{1, 16}))
	fmt.Println(insert([][]int{{1, 5}}, []int{6, 8}))
	fmt.Println(insert([][]int{{1, 5}}, []int{0, 0}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
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
