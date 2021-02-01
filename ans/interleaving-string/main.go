package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if s1 == "" {
		return s2 == s3
	}
	if s2 == "" {
		return s1 == s3
	}
	m := make([][]byte, len(s1))
	for i := range m {
		m[i] = make([]byte, len(s2))
	}
	return isi(m, s1, s2, s3, len(s1)-1, len(s2)-1)
}

func isi(m [][]byte, s1, s2, s3 string, l1, l2 int) (ok bool) {
	if l1 < 0 {
		return s3[:l2+1] == s2[:l2+1]
	}
	if l2 < 0 {
		return s3[:l1+1] == s1[:l1+1]
	}
	if m[l1][l2] != 0 {
		return m[l1][l2] == 1
	}
	defer func() {
		if ok {
			m[l1][l2] = 1
		} else {
			m[l1][l2] = 2
		}
	}()
	c1, c2, c3 := s1[l1], s2[l2], s3[l1+l2+1]
	if c1 == c3 {
		if isi(m, s1, s2, s3, l1-1, l2) {
			return true
		}
	}
	if c2 == c3 {
		if isi(m, s1, s2, s3, l1, l2-1) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}
