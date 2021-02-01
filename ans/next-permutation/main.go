package main

import (
	"fmt"
	"sort"
)

func frac(i int) int {
	if i <= 1 {
		return 1
	}
	return i * frac(i-1)
}

func main() {
	for _, a := range [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{3, 2, 1},
		{1},
		{3, 2, 5, 4, 1},
	} {
		fmt.Println(a)
		nextPermutation(a)
		fmt.Println(a)
	}
	a := []int{1, 1, 5}
	for i := 0; i < frac(len(a)); i++ {
		fmt.Println(a)
		nextPermutation(a)
	}
	fmt.Println(a)
}

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}
	if i == 0 {
		sort.Ints(nums)
		return
	}
	sort.Ints(nums[i:])
	v := nums[i-1]
	j := len(nums) - 1
	k := sort.Search(j-i, func(j int) bool { return nums[j+i] > v })
	nums[i-1], nums[k+i] = nums[k+i], nums[i-1]
}
