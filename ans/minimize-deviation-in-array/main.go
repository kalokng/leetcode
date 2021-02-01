package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumDeviation([]int{1, 5, 9}))
	fmt.Println(minimumDeviation([]int{2, 10, 8}))
	fmt.Println(minimumDeviation([]int{1, 2, 3, 4}))
	fmt.Println(minimumDeviation([]int{3, 5}))
	fmt.Println(minimumDeviation([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(minimumDeviation([]int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8}))
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumDeviation(nums []int) int {
	var mn = math.MaxInt32
	for i, n := range nums {
		if n%2 == 1 {
			n *= 2
			nums[i] = n
		}
		mn = min(mn, n)
	}
	h := IntHeap(nums)
	heap.Init(&h)
	r := h[0] - mn
	l := h[0] + 1
	for {
		v := h[0]
		if l == v {
			heap.Pop(&h)
			continue
		}
		r = min(r, v-mn)
		if v%2 != 0 {
			break
		}
		l = v
		v /= 2
		mn = min(mn, v)
		h[0] = v
		heap.Fix(&h, 0)
	}
	return r
}
