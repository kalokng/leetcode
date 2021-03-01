// +build ignore

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(searchMatrix([][]int{
		{1},
	}, 0))
	fmt.Println(searchMatrix([][]int{
		{1},
	}, 1))
	fmt.Println(searchMatrix([][]int{
		{0},
	}, 1))
	fmt.Println(searchMatrix([][]int{
		{1, 3},
	}, 3))
	fmt.Println(searchMatrix([][]int{
		{1},
		{3},
	}, 3))
	return
	m := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println()
	for i := 0; i <= 31; i++ {
		fmt.Println(i, searchMatrix(m, i))
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	mx := sort.Search(m, func(i int) bool {
		return matrix[i][0] > target
	})
	mn := sort.Search(m, func(i int) bool {
		return matrix[i][n-1] >= target
	})
	if mx <= 0 {
		return false
	}
	mx--
	if mn >= m {
		return false
	}
	for i := mn; i <= mx; i++ {
		j := sort.Search(n, func(j int) bool {
			return matrix[i][j] >= target
		})
		if j < n && matrix[i][j] == target {
			return true
		}
	}
	return false
}
