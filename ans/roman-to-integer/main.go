package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

var r = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}
var r2 = map[byte]byte{
	'V': 'I',
	'X': 'I',
	'L': 'X',
	'C': 'X',
	'D': 'C',
	'M': 'C',
}

func romanToInt(s string) int {
	n := 0
	var l byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		n += r[c]
		if t, ok := r2[c]; ok && t == l {
			n -= r[t] * 2
		}
		l = c
	}
	return n
}
