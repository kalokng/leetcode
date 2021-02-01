package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}

	size := len(words[0])
	sl := size * len(words)
	l := len(s) - sl + 1

	var res []int
	for i := 0; i < l; i++ {
		if match(s[i:i+sl], m, size) {
			res = append(res, i)
		}
	}
	return res
}

func match(s string, m map[string]int, l int) bool {
	h := map[string]int{}
	for i := 0; i < len(s); i += l {
		w := s[i : i+l]
		if m[w] == h[w] {
			return false
		}
		h[w]++
	}
	for k, v := range h {
		if m[k] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"bar", "foo"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring("aaaaaa", []string{"a", "a", "a"}))
}
