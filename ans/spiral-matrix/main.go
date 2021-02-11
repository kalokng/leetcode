package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}}))
	fmt.Println(spiralOrder([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}}))
	fmt.Println(spiralOrder([][]int{{1}, {2}, {3}, {4}, {5}, {6}}))
	fmt.Println(spiralOrder([][]int{{1}}))
}

func spiralOrder(matrix [][]int) []int {
	return so(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1)
}

func so(m [][]int, t, b, l, r int) []int {
	if b-t == 0 {
		return m[t][l : r+1]
	}
	var res = m[t][l : r+1]
	for i := t + 1; i < b; i++ {
		res = append(res, m[i][r])
	}
	for i := r; i >= l; i-- {
		res = append(res, m[b][i])
	}
	if l < r {
		for i := b - 1; i > t; i-- {
			res = append(res, m[i][l])
		}
	}
	if b-t <= 1 || r-l <= 1 {
		return res
	}
	return append(res, so(m, t+1, b-1, l+1, r-1)...)
}
