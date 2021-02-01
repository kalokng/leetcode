package main

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	s = strings.Replace(s, "E", "e", -1)
	sub := strings.SplitN(s, "e", 2)
	if len(sub) > 1 && !isInteger(sub[1]) {
		return false
	}
	return isDecimal(sub[0])
}

func isDecimal(s string) bool {
	if s == "" {
		return false
	}
	c := s[0]
	if c == '+' || c == '-' {
		s = s[1:]
	}
	dot := -1
	for i, c := range s {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case '.':
			if dot != -1 {
				return false
			}
			dot = i
		default:
			return false
		}
	}
	if dot < 0 {
		return len(s) > 0
	}
	return s != "."
}

func isInteger(s string) bool {
	if s == "" {
		return false
	}
	c := s[0]
	if c == '+' || c == '-' {
		s = s[1:]
	}
	for _, c := range s {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			return false
		}
	}
	return len(s) > 0
}

func main() {
	for _, s := range []string{
		"2",
		"0089",
		"-0.1",
		"+3.14",
		"4.",
		"-.9",
		"2e10",
		"-90E3",
		"3e+7",
		"+6e-1",
		"53.5e93",
		"-123.456e789",
		"abc",
		"1a",
		"1e",
		"e3",
		"99e2.5",
		"--6",
		"-+3",
		"95a54e53",
	} {
		fmt.Println(s, isNumber(s))
	}
}
