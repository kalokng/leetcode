package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	fmt.Println(firstMissingPositive([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(firstMissingPositive([]int{6, 5, 4, 3, 2, 1}))
	fmt.Println(firstMissingPositive([]int{3, 1, 4, 1, 5, 9, 2}))
}

func firstMissingPositive(nums []int) int {
loop:
	for i := len(nums) - 1; i >= 0; i-- {
		v := nums[i] - 1
		for v != i {
			if v < 0 || v > i || v == nums[v]-1 {
				nums = nums[:i]
				continue loop
			}
			nums[v], nums[i] = nums[i], nums[v]
			v = nums[i] - 1
		}
	}
	return len(nums) + 1
}
