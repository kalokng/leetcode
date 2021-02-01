package main

import "fmt"

var sq = []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}

func isHappy(n int) bool {
	m := map[int]bool{n: true}
	for n != 1 {
		var sum int
		for n >= 10 {
			sum, n = sum+sq[n%10], n/10
		}
		n = sum + sq[n]
		if m[n] {
			return false
		}
		m[n] = true
	}
	return true
}

func main() {
	fmt.Println(isHappy(1<<31 - 1))
}
