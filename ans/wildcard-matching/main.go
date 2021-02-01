package main

import "fmt"

type matched struct {
	m    []byte
	lenp int
	s, p string
}

func (m *matched) get(s, p int) (hit bool, ok bool) {
	i := s*m.lenp + p
	if i >= len(m.m) {
		return false, false
	}
	b := m.m[i]
	return b == 1, b != 0
}
func (m *matched) set(s, p int, ok bool) {
	i := s*m.lenp + p
	if i >= len(m.m) {
		return
	}
	var b byte = 2
	if ok {
		b = 1
	}
	m.m[i] = b
}

func isMatch(s string, p string) bool {
	m := matched{
		m:    make([]byte, (len(s)+1)*len(p)),
		lenp: len(p),
		s:    s,
		p:    p,
	}
	return m.ismatch(len(s), len(p))
}

func (m *matched) ismatch(ls, lp int) (ok bool) {
	if b, ok := m.get(ls, lp); ok {
		return b
	}

	defer func() {
		m.set(ls, lp, ok)
	}()

	if lp == 0 {
		return ls == 0
	}
	if ls == 0 {
		for i := 0; i < lp; i++ {
			if m.p[i] != '*' {
				return false
			}
		}
		return true
	}
	ep := m.p[lp-1]
	switch ep {
	case '?':
		return m.ismatch(ls-1, lp-1)
	case '*':
		for i := 0; i <= ls; i++ {
			if m.ismatch(i, lp-1) {
				return true
			}
		}
		return false
	}
	if ep == m.s[ls-1] {
		return m.ismatch(ls-1, lp-1)
	}
	return false
}

func main() {
	test := func(s, p string) {
		fmt.Println(s, p, isMatch(s, p))
	}
	//test("aa", "a")
	//test("aa", "*")
	//test("cb", "?a")
	//test("adceb", "*a*b")
	//test("acdcb", "a*c?b")
	//test("acdcb", "***")
	//test("", "")
	//test("", "*")
	//test("", "?")
	test("acdcb", "a*c*b")
}
