package main

import "fmt"

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 1, 2, 5, 7}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4, 3, 2, 1}))
}

func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}
	i := 1
	d := A[1] - A[0]
	n := 0
	for j := 2; j < len(A); j++ {
		nd := A[j] - A[j-1]
		if d != nd {
			i = j
			d = nd
			continue
		}
		n += j - i
	}
	return n
}
