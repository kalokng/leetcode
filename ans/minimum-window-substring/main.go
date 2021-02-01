package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	m := make(map[byte]int)
	h := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		v := t[i]
		n := h[v]
		h[v] = n + 1
	}
	var l, r int
	var ok int
	for r = 0; r < len(s); r++ {
		c := s[r]
		if h[c] == 0 {
			continue
		}
		m[c]++
		if m[c] == h[c] {
			ok++
			if ok == len(h) {
				break
			}
		}
	}
	if ok < len(h) {
		return ""
	}
	var min = 100000
	var i = 0
	for r < len(s) {
		for ; l <= r; l++ {
			c := s[l]
			if h[c] == 0 {
				continue
			}
			if m[c] == h[c] {
				break
			}
			m[c]--
		}
		if min > r-l {
			min = r - l
			i = l
		}
		r++
		for ; r < len(s); r++ {
			c := s[r]
			if h[c] == 0 {
				continue
			}
			m[c]++
			if s[r] == s[l] {
				break
			}
		}
	}
	return s[i : i+min+1]
}

func main() {
	fmt.Println(minWindow("bba", "ab"))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("ADOBECODEBANC", "AABC"))
}
