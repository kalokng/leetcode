package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3/2 "))
	fmt.Println(calculate(" 3+5 / 2 "))
	fmt.Println(calculate(" 3 + 5 / 2 "))
	fmt.Println(calculate(" 31/2 "))
	fmt.Println(calculate(" 31+5 / 2 "))
	fmt.Println(calculate(" 31 + 5 / 2 "))
	fmt.Println(calculate(" 0 + 5 / 2 "))
	fmt.Println(calculate(" 123 + 5 / 2 "))
	fmt.Println(calculate(" 1+2+3+4*5 "))
	fmt.Println(calculate("1*2-3/4+5*6-7*8+9/10"))
}

func isdigit(c rune) bool {
	return c >= '0' && c <= '9'
}
func notdigit(c rune) bool {
	return c < '0' || c > '9'
}
func readInt(s string) (int, string) {
	i := strings.IndexFunc(s, isdigit)
	j := strings.IndexFunc(s[i:], notdigit)
	if j < 0 {
		j = len(s)
	} else {
		j += i
	}
	n := 0
	for i < j {
		n = n*10 + int(s[i]-'0')
		i++
	}
	return n, s[i:]
}
func readOp(s string) (byte, string) {
	i := strings.IndexAny(s, "+-*/")
	if i == -1 {
		return 0, s
	}
	return s[i], s[i+1:]
}

func calMul(s string) (int, string) {
	v, s := readInt(s)
	var op byte
	for {
		os := s
		op, s = readOp(s)
		switch op {
		case 0:
			return v, s
		case '*':
			var v2 int
			v2, s = readInt(s)
			v = v * v2
		case '/':
			var v2 int
			v2, s = readInt(s)
			v = v / v2
		case '+', '-':
			return v, os
		}
	}
}

func calculate(s string) int {
	v, s := readInt(s)
	for {
		var op byte
		op, s = readOp(s)
		switch op {
		case 0:
			return v
		case '*':
			var v2 int
			v2, s = readInt(s)
			v = v * v2
		case '/':
			var v2 int
			v2, s = readInt(s)
			v = v / v2
		case '+':
			var v2 int
			op, _ = readOp(s)
			if op != '+' && op != '-' {
				v2, s = calMul(s)
			} else {
				v2, s = readInt(s)
			}
			v = v + v2
		case '-':
			var v2 int
			op, _ = readOp(s)
			if op != '+' && op != '-' {
				v2, s = calMul(s)
			} else {
				v2, s = readInt(s)
			}
			v = v - v2
		}
	}
}
