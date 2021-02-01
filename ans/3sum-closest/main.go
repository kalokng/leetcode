package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 0, 1, 2, -1, -4}, 1))
	fmt.Println(threeSumClosest([]int{-1, 0, 1, 2, -1, -4}, 1))
	fmt.Println(threeSumClosest([]int{-1, -1, 0, 1, 2, -1, -4}, 1))
	fmt.Println(threeSumClosest([]int{3, 0, -2, -1, 1, 2}, 1))
	fmt.Println(threeSumClosest([]int{-2, 0, 0, 2, 2}, 1))
	fmt.Println(threeSumClosest([]int{-1, 2, 1, 4}, 1))
}

func threeSumClosest(nums []int, target int) int {
	var d int = 100000
	var res int
	sort.Ints(nums)
	last := nums[0] - 1
	for idx, v := range nums {
		if last == v {
			continue
		}
		i, j := idx+1, len(nums)-1
		for i < j {
			r := nums[i] + nums[j] + v
			a := abs(r - target)
			if a < d {
				res = r
				d = a
			}
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
				return target
			}
		}
		last = v
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
