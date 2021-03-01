package main

import "fmt"

func main() {
	fmt.Println(scoreOfParentheses("()"))
	fmt.Println(scoreOfParentheses("(())"))
	fmt.Println(scoreOfParentheses("()()"))
	fmt.Println(scoreOfParentheses("(()(()))"))
	fmt.Println(scoreOfParentheses("()()()"))
	fmt.Println(scoreOfParentheses("()(())()"))
}

func scoreOfParentheses(S string) int {
	var n = []int{0}
	for i := 0; i < len(S); i++ {
		c := S[i]
		switch c {
		case '(':
			n = append(n, 0)
		case ')':
			if n[len(n)-1] == 0 {
				n = n[:len(n)-1]
				n[len(n)-1]++
				break
			}
			n[len(n)-2] += n[len(n)-1] * 2
			n = n[:len(n)-1]
		}
	}
	return n[0]
}
