package main

import (
	"fmt"
	"sort"
)

func main() {
	m := [][]int{
		{5, 2},
		{1, 6},
	}
	fmt.Println(kthLargestValue(m, 1))
	fmt.Println(kthLargestValue(m, 2))
	fmt.Println(kthLargestValue(m, 3))
	fmt.Println(kthLargestValue(m, 4))
	m = [][]int{
		{10, 9, 5},
		{2, 0, 4},
		{1, 0, 9},
		{3, 4, 8},
	}
	fmt.Println(kthLargestValue(m, 10))
}

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	xor := make([]int, m*n)
	fill(xor, m, n, matrix)
	sort.Ints(xor)
	return xor[m*n-k]
}

func fill(xor []int, m, n int, matrix [][]int) {
	xor[0] = matrix[0][0]
	for i := 1; i < m; i++ {
		xor[i*n] = xor[(i-1)*n] ^ matrix[i][0]
	}
	for i := 1; i < n; i++ {
		xor[i] = xor[i-1] ^ matrix[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			v := i*n + j
			xor[v] = xor[v-1] ^ xor[v-n] ^ xor[v-n-1] ^ matrix[i][j]
		}
	}
}
