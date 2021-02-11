// +build ignore

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
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	i := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][0] >= newInterval[0]
	})
	j := i
	if i >= len(intervals) {
		i--
	}
	for i >= 0 {
		if intervals[i][1] < newInterval[0] {
			break
		}
		i--
	}
	i++
	if i < len(intervals) && newInterval[0] <= intervals[i][1] && newInterval[1] >= intervals[i][0] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
	}
	for j < len(intervals) {
		if intervals[j][0] > newInterval[1] {
			break
		}
		j++
	}
	if j > 0 && newInterval[0] <= intervals[j-1][1] && newInterval[1] >= intervals[j-1][0] {
		newInterval[0] = min(newInterval[0], intervals[j-1][0])
		newInterval[1] = max(newInterval[1], intervals[j-1][1])
	}
	return append(intervals[:i], append([][]int{newInterval}, intervals[j:]...)...)
}
