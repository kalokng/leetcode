package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minCharacters("aba", "caa"))
	fmt.Println(minCharacters("dabadd", "cda"))
	fmt.Println(minCharacters("bababd", "cda"))
	fmt.Println(minCharacters("a", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"))
}

func minCharacters(a string, b string) int {
	var ma = make([]int, 26)
	var mb = make([]int, 26)
	var mc = make([]int, 26)
	for i := 0; i < len(a); i++ {
		v := a[i] - 'a'
		mc[v]++
		ma[v]++
	}
	for i := 0; i < len(b); i++ {
		v := b[i] - 'a'
		mc[v]++
		mb[v]++
	}
	sort.Ints(mc)
	mna := len(a) + mincount(ma, mb)
	mnb := len(b) + mincount(mb, ma)
	mnc := len(a) + len(b) - mc[len(mc)-1]

	mn := min(mna, mnb)
	mn = min(mn, mnc)
	return mn
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func mincount(a, b []int) int {
	na := a[0]
	nb := b[0]
	mn := nb - na
	for i := 1; i < len(a)-1; i++ {
		na += a[i]
		nb += b[i]
		mn = min(mn, nb-na)
	}
	return mn
}
