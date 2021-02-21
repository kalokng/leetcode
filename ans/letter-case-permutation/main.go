package main

import "fmt"

func main() {
	fmt.Println(letterCasePermutation("2b9s"))
}

func letterCasePermutation(S string) []string {
	return lp(S, 0)
}

func lp(s string, i int) []string {
	if i >= len(s) {
		return []string{s}
	}
	c := s[i]
	if c >= '0' && c <= '9' {
		return lp(s, i+1)
	}
	b := []byte(s)
	if c >= 'a' && c <= 'z' {
		b[i] = c - 'a' + 'A'
	} else {
		b[i] = c - 'A' + 'a'
	}
	return append(lp(s, i+1), lp(string(b), i+1)...)
}
