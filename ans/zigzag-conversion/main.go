package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var b strings.Builder
	p := 2*numRows - 2
	b.Write(sel(s, p, 0))
	for i := 1; i < numRows-1; i++ {
		b.Write(sel(s, p, i, p-i))
	}
	b.Write(sel(s, p, numRows-1))
	return b.String()
}

func sel(s string, p int, idx ...int) []byte {
	b := make([]byte, 0, (len(s)/p+1)*len(idx))
	for i := 0; i < len(s); i += p {
		for _, j := range idx {
			if i+j >= len(s) {
				break
			}
			b = append(b, s[i+j])
		}
	}
	return b
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 1))
	fmt.Println(convert("PAYPALISHIRING", 2))
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 5))
}
