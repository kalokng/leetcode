package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minTrioDegree(6, [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}))
	fmt.Println(minTrioDegree(7, [][]int{{1, 3}, {4, 1}, {4, 3}, {2, 5}, {5, 6}, {6, 7}, {7, 5}, {2, 6}}))
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minTrioDegree(n int, edges [][]int) int {
	m := make([][]bool, n)
	for i := 0; i < n; i++ {
		m[i] = make([]bool, n)
	}
	c := make([]int, n)

	for _, v := range edges {
		a, b := v[0]-1, v[1]-1
		if a > b {
			a, b = b, a
		}
		c[a]++
		c[b]++
		m[a][b] = true
	}
	mn := math.MaxInt32
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !m[i][j] {
				continue
			}
			sum := c[i] + c[j]
			if sum >= mn {
				continue
			}
			for k := j + 1; k < n; k++ {
				if !m[i][k] || !m[j][k] {
					continue
				}
				mn = min(mn, sum+c[k])
			}
		}
	}
	if mn == math.MaxInt32 {
		return -1
	}
	return mn - 6
}
