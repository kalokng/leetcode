package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return iss(s1, s2, 0, len(s1), 0, len(s2))
}

func iss(s1, s2 string, l1, r1, l2, r2 int) bool {
	if r1-l1 == 1 {
		return s1[l1:r1] == s2[l2:r2]
	}
}

func main() {
	fmt.Println(isScramble("great", "rgeat"))
}
