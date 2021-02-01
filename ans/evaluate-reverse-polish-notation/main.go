package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

var op = map[string]func(a, b int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func mustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func evalRPN(tokens []string) int {
	var st []int
	for _, t := range tokens {
		f, ok := op[t]
		if !ok {
			st = append(st, mustAtoi(t))
			continue
		}
		n := len(st) - 1
		a, b := st[n-1], st[n]
		st = append(st[:n-1], f(a, b))
	}
	return st[0]
}
