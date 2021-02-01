package main

import (
	"fmt"
	"sort"
)

func main() {
	m := [][][]int{
		{
			{3, 3, 1, 1},
			{2, 2, 1, 2},
			{1, 1, 1, 2},
		},
		{
			{3, 3, 1, 1},
			{2, 2, 1, 2},
			{1, 1, 1, 2},
			{4, 5, 1, 2},
			{1, 6, 7, 2},
			{8, 9, 10, 12},
			{1, 12, 13, 2},
		},
	}
	for _, x := range m {
		diagonalSort(x)
		fmt.Println(x)
	}
}

type diagonal struct {
	mat  [][]int
	r, c int
	x, y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (d diagonal) Len() int {
	return min(d.r-d.y, d.c-d.x)
}

func (d diagonal) Less(i, j int) bool {
	return d.mat[i+d.y][i+d.x] < d.mat[j+d.y][j+d.x]
}

func (d diagonal) Swap(i, j int) {
	d.mat[i+d.y][i+d.x], d.mat[j+d.y][j+d.x] = d.mat[j+d.y][j+d.x], d.mat[i+d.y][i+d.x]
}

func diagonalSort(mat [][]int) [][]int {
	if len(mat) == 0 {
		return nil
	}
	r, c := len(mat), len(mat[0])
	d := diagonal{mat: mat, r: r, c: c, x: 0}
	for i := 0; i < r; i++ {
		d.y = i
		sort.Sort(d)
	}
	d.y = 0
	for j := 1; j < c; j++ {
		d.x = j
		sort.Sort(d)
	}
	return mat
}
