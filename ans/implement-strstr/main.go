package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("hello", ""))
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("aabba", "bba"))
	fmt.Println(strStr("a", "a"))
}

func strStr(haystack string, needle string) int {
	l := len(needle)
	if l == 0 {
		return 0
	}
	for i := l; i <= len(haystack); i++ {
		if haystack[i-l:i] == needle {
			return i - l
		}
	}
	return -1
}
