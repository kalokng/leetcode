package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{1, 1}))
	fmt.Println(findDuplicate([]int{1, 1, 2}))
	fmt.Println(findDuplicate([]int{2, 2, 2, 2, 2}))
	fmt.Println(findDuplicate([]int{1, 4, 4, 2, 4}))
	fmt.Println(findDuplicate([]int{4, 4, 4, 4, 4, 4, 4}))
	fmt.Println(findDuplicate([]int{1, 6, 2, 5, 4, 4, 4}))
}

func findDuplicate(nums []int) int {
	t := nums[0]
	h := nums[t]
	for t != h {
		t = nums[t]
		h = nums[nums[h]]
	}
	t = 0
	for t != h {
		t = nums[t]
		h = nums[h]
	}
	return t
}
