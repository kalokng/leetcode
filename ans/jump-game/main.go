package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{2, 3, 0, 1, 4}))
	fmt.Println(canJump([]int{2, 3, 1, 0, 4}))
	fmt.Println(canJump([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	reach := 0
	for i := 0; i < len(nums); i++ {
		reach = max(reach, nums[i]+i)
		if reach >= len(nums)-1 {
			return true
		}
		if i == reach {
			return false
		}
	}
	return true
}
