package main

import "fmt"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	maxi := 0
	max := height[0]
	for i, v := range height {
		if v > max {
			max = v
			maxi = i
		}
	}
	max = 0
	sum := 0
	for _, v := range height[:maxi] {
		if v > max {
			max = v
			continue
		}
		sum += max - v
	}
	max = 0
	for i := len(height) - 1; i > maxi; i-- {
		v := height[i]
		if v > max {
			max = v
			continue
		}
		sum += max - v
	}
	return sum
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
