package main

import "fmt"

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))
	fmt.Println(findLongestWord("aaa", []string{"aaa", "aa", "a"}))
}

func findLongestWord(s string, d []string) string {
	if len(s) == 0 {
		return ""
	}
	h := make([][]int, 26)
	for i := range h {
		h[i] = make([]int, len(s))
		h[i][len(s)-1] = -1
	}
	c := s[len(s)-1]
	h[c-'a'][len(s)-1] = len(s) - 1
	for i := len(s) - 2; i >= 0; i-- {
		c = s[i]
		for j := range h {
			h[j][i] = h[j][i+1]
		}
		h[c-'a'][i] = i
	}

	var mx string
	for _, v := range d {
		j := 0
		for i := 0; i < len(v); i++ {
			if j >= len(s) {
				j = -1
				break
			}
			c := v[i]
			j = h[c-'a'][j]
			if j == -1 {
				break
			}
			j++
		}
		if j == -1 {
			continue
		}
		if len(v) > len(mx) || (len(v) == len(mx) && mx > v) {
			mx = v
		}
	}
	return mx
}
