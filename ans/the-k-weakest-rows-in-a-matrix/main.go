package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}, 3))
	fmt.Println(kWeakestRows([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}, 2))
}

func kWeakestRows(mat [][]int, k int) []int {
	for i, r := range mat {
		s := 0
		for _, c := range r {
			s += c
		}
		r[0] = s
		r[1] = i
	}
	sort.SliceStable(mat, func(i, j int) bool {
		return mat[i][0] < mat[j][0]
	})
	rank := make([]int, k)
	for i := range rank {
		rank[i] = mat[i][1]
	}
	return rank
}
