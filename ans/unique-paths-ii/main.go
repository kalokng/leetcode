package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 1},
		{0, 0},
	}))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	r, c := len(obstacleGrid), len(obstacleGrid[0])
	obstacleGrid[0][0] = 1
	for i := 1; i < r; i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
			continue
		}
		obstacleGrid[i][0] = obstacleGrid[i-1][0]
	}
	for j := 1; j < c; j++ {
		if obstacleGrid[0][j] == 1 {
			obstacleGrid[0][j] = 0
			continue
		}
		obstacleGrid[0][j] = obstacleGrid[0][j-1]
	}
	for s := 1; s < r+c; s++ {
		i := min(r-1, s)
		j := s - i
		if j == 0 {
			i--
			j++
		}
		for j < c && i > 0 {
			fmt.Println(i, j)
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				i--
				j++
				continue
			}
			obstacleGrid[i][j] += obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			i--
			j++
		}
	}
	return obstacleGrid[r-1][c-1]
}
