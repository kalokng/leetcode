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
	m := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println()
	for i := 0; i <= 61; i++ {
		fmt.Println(i, searchMatrix(m, i))
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	i := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	})
	if i <= 0 {
		return false
	}
	i--
	j := sort.Search(len(matrix[i]), func(j int) bool {
		return matrix[i][j] >= target
	})
	if j >= len(matrix[i]) {
		return false
	}
	return matrix[i][j] == target
}
