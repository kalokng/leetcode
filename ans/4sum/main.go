package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{-1, 0, 1, 2, -1, -4}, 0))
	fmt.Println(fourSum([]int{-1, 0, 1, 2, -1, -4}, 1))
	fmt.Println(fourSum([]int{-1, -1, 0, 1, 2, -1, -4}, 2))
	fmt.Println(fourSum([]int{3, 0, -2, -1, 1, 2}, 3))
	fmt.Println(fourSum([]int{-2, -2, 0, 0, 2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
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
		jlast := v - 1
		for jdx := idx + 1; jdx < len(nums); jdx++ {
			vv := nums[jdx]
			if jlast == vv {
				continue
			}
			vvv := v + vv
			i, j := jdx+1, len(nums)-1
			for i < j {
				r := nums[i] + nums[j] + vvv
				switch {
				case r < target:
					i++
					for i < j && nums[i] == nums[i-1] {
						i++
					}
				case r > target:
					j--
					for i < j && nums[j] == nums[j+1] {
						j--
					}
				default:
					res = append(res, []int{v, vv, nums[i], nums[j]})
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
			jlast = vv
		}
		last = v
	}
	return res
}
