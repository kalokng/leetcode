package main

import "fmt"

func main() {
	fmt.Println(kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 3))
	fmt.Println(kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2))
	fmt.Println(kLengthApart([]int{1, 0, 0, 0, 1, 0, 1}, 2))
	fmt.Println(kLengthApart([]int{1, 1, 1, 1, 1, 1, 1}, 0))
	fmt.Println(kLengthApart([]int{0, 1, 0, 1}, 1))
	fmt.Println(kLengthApart([]int{0, 1, 0, 1}, 2))
}

func kLengthApart(nums []int, k int) bool {
	d := 0
	for _, v := range nums {
		if v == 0 {
			d--
			continue
		}
		if d > 0 {
			return false
		}
		d = k
	}
	return true
}
