package main

import "fmt"

func longestValidParentheses(s string) int {
	matched := make([]bool, len(s))
	m := make(map[int]int)

	opened := 0
	for i, c := range s {
		switch c {
		case '(':
			m[opened] = i
			opened++
		case ')':
			if opened == 0 {
				break
			}
			opened--
			matched[m[opened]] = true
			matched[i] = true
		}
	}
	max := 0
	n := 0
	for _, ok := range matched {
		if !ok {
			n = 0
			continue
		}
		n++
		if n > max {
			max = n
		}
	}
	return max
}

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(""))
	fmt.Println(longestValidParentheses("())()()"))
	fmt.Println(longestValidParentheses("(((()"))
	fmt.Println(longestValidParentheses("(((((((()"))
	fmt.Println(longestValidParentheses("()(()"))
}
