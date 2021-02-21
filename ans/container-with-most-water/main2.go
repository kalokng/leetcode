package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
	fmt.Println(maxArea([]int{1, 2, 1}))
	fmt.Println(maxArea([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(maxArea([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(maxArea([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(maxArea([]int{2, 4, 6, 8, 10, 9, 7, 5, 3, 1}))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxArea(height []int) int {
	var area int
	for len(height) > 0 {
		l := len(height) - 1
		if height[0] < height[l] {
			area = max(area, height[0]*(len(height)-1))
			height = height[1:]
			continue
		}
		area = max(area, height[l]*l)
		height = height[:l]
	}
	return area
}
