package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
	fmt.Println(maxArea([]int{1, 2, 1}))
	fmt.Println(maxArea([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(maxArea([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(maxArea([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(maxArea([]int{2, 4, 6, 8, 10, 9, 7, 5, 3, 1}))
}

type pairHeap struct {
	h []int
	i []int
}

func (h pairHeap) Len() int           { return len(h.i) }
func (h pairHeap) Less(i, j int) bool { return h.h[h.i[i]] > h.h[h.i[j]] }
func (h pairHeap) Swap(i, j int)      { h.i[i], h.i[j] = h.i[j], h.i[i] }
func (h *pairHeap) Push(x interface{}) {
	h.i = append(h.i, x.(int))
}
func (h *pairHeap) Pop() interface{} {
	old := h.i
	n := len(old)
	x := old[n-1]
	h.i = old[:n-1]
	return x
}

func maxArea(height []int) int {
	var p = &pairHeap{
		h: height,
		i: make([]int, 0, len(height)),
	}
	for i := range height {
		p.i = append(p.i, i)
	}

	heap.Init(p)
	i := heap.Pop(p).(int)
	j := i
	h := height[j]
	var max int
	for len(p.i) > 0 {
		k := heap.Pop(p).(int)
		switch {
		case k < i:
			h = height[k]
			i = k
			if h*(j-i) > max {
				max = h * (j - i)
			}
		case k > j:
			h = height[k]
			j = k
			if h*(j-i) > max {
				max = h * (j - i)
			}
		}
	}
	return max
}
