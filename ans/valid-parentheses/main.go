package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("}{"))
	fmt.Println(isValid("([{}])"))
	fmt.Println(isValid("("))
}

var pair = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		default:
			n := len(stack) - 1
			if n < 0 || stack[n] != pair[c] {
				return false
			}
			stack = stack[:n]
		}
	}
	return len(stack) == 0
}
