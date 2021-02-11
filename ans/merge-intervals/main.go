// +build ignore

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 3}, {4, 10}, {3, 4}, {15, 18}}))
}

type pHeap [][]int

func (h pHeap) Len() int            { return len(h) }
func (h pHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h pHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *pHeap) Push(x interface{}) {}
func (h *pHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	var h pHeap = intervals
	heap.Init(&h)
	var res = [][]int{intervals[0]}
	heap.Pop(&h)
	for len(h) > 0 {
		p := h[0]
		heap.Pop(&h)
		pp := res[len(res)-1]
		if pp[1] >= p[0] {
			pp[1] = max(pp[1], p[1])
			continue
		}
		res = append(res, p)
	}
	return res
}
