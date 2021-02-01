package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-1, -1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{3, 0, -2, -1, 1, 2}))
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var res [][]int
	sort.Ints(nums)
	last := nums[0] - 1
	for idx, v := range nums {
		if last == v {
			continue
		}
		i, j := idx+1, len(nums)-1
		for i < j {
			r := nums[i] + nums[j] + v
			switch {
			case r < 0:
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			case r > 0:
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			default:
				res = append(res, []int{v, nums[i], nums[j]})
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}
		last = v
	}
	return res
}
