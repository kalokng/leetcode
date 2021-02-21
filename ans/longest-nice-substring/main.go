package main

import "fmt"

func main() {
	fmt.Println(longestNiceSubstring("YazaAay"))
	fmt.Println(longestNiceSubstring("qQUjJ"))
	fmt.Println(longestNiceSubstring("zUXxizubXNKAUGXTjmAXkpzNZMnRBgddDUAWPUa"))
}

func nice(s string) bool {
	m := map[byte]struct{}{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		m[c] = struct{}{}
	}
	for k := range m {
		if k > 'z' || k < 'a' {
			if _, ok := m[k-'A'+'a']; !ok {
				return false
			}
			continue
		}
		if _, ok := m[k-'a'+'A']; !ok {
			return false
		}
	}
	return true
}

func longestNiceSubstring(s string) string {
	mx := 0
	imx := 0
	i := 0
	for ; i < len(s)-mx; i++ {
		for j := i + 1 + mx; j <= len(s); j++ {
			if nice(s[i:j]) {
				mx = j - i
				imx = i
			}
		}
	}
	return s[imx : imx+mx]
}
