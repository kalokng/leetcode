package main

import "fmt"

func main() {
	for _, s := range []string{
		"babad",
		"x",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaa",
	} {
		a := longestPalindrome(s)
		fmt.Println(len(a), a)
	}
}

func longestPalindrome(s string) string {
	var max int
	var l, r int
	for i := 2*len(s) - 2; i >= 0; i-- {
		v := pal(s, i)
		if v > max {
			max = v
			v = v - 1
			l = (i - v) / 2
			r = (i + v) / 2
		}
	}
	return s[l : r+1]
}

func pal(s string, p int) (l int) {
	i := p / 2
	j := p - i
	eq := 0
	if i == j {
		eq = 1
		i--
		j++
	}
	for i >= 0 && j < len(s) && s[i] == s[j] {
		eq += 2
		i--
		j++
	}
	return eq
}
