package main

import "fmt"

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{2, 5, 6, 4, 8, 11, 9, 10, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{1}))
	fmt.Println(findUnsortedSubarray([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}))
	fmt.Println(findUnsortedSubarray([]int{2, 4, 6, 13, 8, 10, 12}))
	fmt.Println(findUnsortedSubarray([]int{2, 4, 6, 1, 8, 10, 12}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 4, 6, 8, 10, 12}))
	fmt.Println(findUnsortedSubarray([]int{2, 1, 4, 6, 8, 10, 12}))
	fmt.Println(findUnsortedSubarray([]int{2, 4, 1, 6, 8, 10, 12}))
	fmt.Println(findUnsortedSubarray([]int{2, 4, 6, 1, 8, 10, 12}))
	fmt.Println(findUnsortedSubarray([]int{2, 4, 6, 8, 1, 10, 12}))
	fmt.Println(findUnsortedSubarray([]int{2, 4, 6, 8, 10, 1, 12}))
	fmt.Println(findUnsortedSubarray([]int{2, 4, 6, 8, 10, 12, 1}))
	fmt.Println(findUnsortedSubarray([]int{1, 3, 2, 3, 3}))
	fmt.Println(findUnsortedSubarray([]int{1, 1, 2, 1, 3}))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findUnsortedSubarray(nums []int) int {
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[start] {
			break
		}
		start++
	}
	if start == len(nums)-1 {
		return 0
	}
	mn := nums[start]
	for i := start + 1; i < len(nums); i++ {
		mn = min(mn, nums[i])
	}
	for i := start; i >= 0; i-- {
		if mn >= nums[i] {
			break
		}
		start--
	}
	start++
	end := len(nums) - 1
	for i := len(nums) - 2; i >= start; i-- {
		if nums[i] > nums[end] {
			break
		}
		end--
	}
	mx := nums[end]
	for i := end - 1; i >= start; i-- {
		mx = max(mx, nums[i])
	}
	for i := end; i < len(nums); i++ {
		if mx <= nums[i] {
			break
		}
		end++
	}
	return end - start
}
