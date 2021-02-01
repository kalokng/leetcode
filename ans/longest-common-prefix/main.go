package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s := strs[0]
	l := len(s)
	for i := 1; i < len(strs); i++ {
		l = min(l, len(strs[i]))
	}
	for j := 0; j < l; j++ {
		c := s[j]
		for i := 1; i < len(strs); i++ {
			if strs[i][j] != c {
				return strs[i][:j]
			}
		}
	}
	return s[:l]
}
