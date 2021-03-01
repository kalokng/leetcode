package main

import (
	"fmt"
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
	i := 0
	j := n - 1
	for i < m && j >= 0 {
		v := matrix[i][j]
		if v == target {
			return true
		}
		if v < target {
			i++
		} else {
			j--
		}
	}
	return false
}
