package main

import (
	"container/heap"
	"fmt"
)

func main() {
	for _, h := range [][][]int{
		{
			{1, 2, 2},
			{3, 8, 2},
			{5, 3, 5},
		},
		{
			{1, 2, 3},
			{3, 8, 4},
			{5, 3, 5},
		},
		{
			{1, 2, 1, 1, 1},
			{1, 2, 1, 2, 1},
			{1, 2, 1, 2, 1},
			{1, 2, 1, 2, 1},
			{1, 1, 1, 2, 1},
		},
		{
			{1, 2, 2, 1, 1},
			{1, 2, 2, 2, 1},
			{1, 2, 2, 2, 1},
			{1, 2, 2, 2, 1},
			{1, 1, 1, 2, 1},
		},
		{
			{1, 2, 2, 1, 1},
			{1, 2, 2, 2, 1},
			{1, 2, 2, 2, 1},
			{1, 2, 2, 2, 1},
			{1, 1, 1, 3, 1},
		},
		{
			{1, 10, 6, 7, 9, 10, 4, 9},
		},
	} {
		fmt.Println(minimumEffortPath(h))
	}
}

func minimumEffortPath(heights [][]int) int {
	mr := len(heights)
	mc := len(heights[0])
	d := make([][]bool, mr)
	for i := range d {
		d[i] = make([]bool, mc)
	}
	return path(d, heights, mr-1, mc-1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type cell struct {
	d, v int
	x, y int
}

type cellHeap []*cell

func (h cellHeap) Len() int { return len(h) }
func (h cellHeap) Less(i, j int) bool {
	return h[i].d < h[j].d
}
func (h cellHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *cellHeap) Push(x interface{}) {
	*h = append(*h, x.(*cell))
}
func (h *cellHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}

func next(ch *cellHeap, d [][]bool, h [][]int, m, mr, mc, x, y int) {
	v := h[x][y]
	app := func(x, y int) {
		if d[x][y] {
			return
		}
		a := abs(v - h[x][y])
		heap.Push(ch, &cell{
			d: a,
			v: max(a, m),
			x: x,
			y: y,
		})
	}
	if x > 0 {
		app(x-1, y)
	}
	if y > 0 {
		app(x, y-1)
	}
	if x < mr {
		app(x+1, y)
	}
	if y < mc {
		app(x, y+1)
	}
}

func path(d [][]bool, h [][]int, mr, mc int) (min int) {
	var ch cellHeap
	next(&ch, d, h, 0, mr, mc, 0, 0)
	for len(ch) > 0 {
		c := ch[0]
		heap.Pop(&ch)
		if d[c.x][c.y] {
			continue
		}
		if c.x == mr && c.y == mc {
			return c.v
		}
		d[c.x][c.y] = true
		next(&ch, d, h, c.v, mr, mc, c.x, c.y)
	}
	return 0
}
