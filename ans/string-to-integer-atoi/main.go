package main

import "fmt"

const max = 1 << 31

func myAtoi(s string) int {
	var i int
	for i = 0; i < len(s); i++ {
		c := s[i]
		if c != ' ' {
			break
		}
	}
	if i == len(s) {
		return 0
	}
	var neg bool
	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		neg = true
	}
	if neg {
		return -atoi(s[i:], max)
	}
	return atoi(s[i:], max-1)
}
func atoi(s string, max int) int {
	var sum int
	for i := 0; i < len(s); i++ {
		d := int(s[i]) - '0'
		if d < 0 || d > 9 {
			return sum
		}
		if sum != 0 && (max-d)/sum < 10 {
			return max
		}
		sum = sum*10 + d
	}
	return sum
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("91283472332"))
}
