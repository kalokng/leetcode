package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := 0
	m := make(map[byte]int)
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if j, ok := m[c]; ok {
			for k := i - n; k < j; k++ {
				delete(m, s[k])
			}
			if n > max {
				max = n
			}
			n = i - j - 1
		}
		m[c] = i
		n++
	}
	if n > max {
		max = n
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("aabacdbeacf"))
}
