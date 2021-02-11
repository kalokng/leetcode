package main

import "fmt"

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
	fmt.Println(shortestToChar("aaab", 'b'))
	fmt.Println(shortestToChar("aaba", 'b'))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func shortestToChar(s string, c byte) []int {
	l := len(s)
	v := make([]int, l)
	d := -l
	for i := 0; i < l; i++ {
		if s[i] == c {
			d = i
		}
		v[i] = i - d
	}
	d = 2 * l
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			d = i
		}
		v[i] = min(v[i], d-i)
	}
	return v
}
