package main

import "fmt"

var fac = []int{
	1,
	1,
	1 * 2,
	1 * 2 * 3,
	1 * 2 * 3 * 4,
	1 * 2 * 3 * 4 * 5,
	1 * 2 * 3 * 4 * 5 * 6,
	1 * 2 * 3 * 4 * 5 * 6 * 7,
	1 * 2 * 3 * 4 * 5 * 6 * 7 * 8,
}

func getPermutation(n int, k int) string {
	var s []byte
	for i := 0; i < n; i++ {
		s = append(s, byte(i+'1'))
	}
	getperm(s, k-1)
	return string(s)
}

func getperm(s []byte, k int) {
	p := len(s) - 1
	for k > 0 {
		for fac[p] > k {
			p--
		}
		var n int
		n, k = k/fac[p], k%fac[p]
		start := len(s) - p - 1
		sel := start + n
		move(s, start, sel)
	}
}

func move(s []byte, start, sel int) {
	v := s[sel]
	for i := sel; i > start; i-- {
		s[i] = s[i-1]
	}
	s[start] = v
}

func main() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(3, 1))

	for i := 1; i <= fac[5]; i++ {
		fmt.Println(getPermutation(5, i))
	}
}
