package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(check([]int{3, 4, 5, 1, 2}))
	fmt.Println(check([]int{2, 1, 3, 4}))
	fmt.Println(check([]int{1, 3, 4}))
	fmt.Println(check([]int{1, 1, 1}))
	fmt.Println(check([]int{2, 1}))
}

func check(nums []int) bool {
	v := nums[len(nums)-1]
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] < v {
			break
		}
		v = nums[i]
	}
	nums = append(nums[i:], nums[:i]...)
	return sort.IsSorted(sort.IntSlice(nums))
}
