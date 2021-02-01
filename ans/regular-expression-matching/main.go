package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(isMatch("aa", "a"))
}

func isMatch(s string, p string) bool {
	m, _ := regexp.MatchString("^"+p+"$", s)
	return m
}
