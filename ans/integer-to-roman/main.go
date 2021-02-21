package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(3999))
}

var r1000 = []string{"", "M", "MM", "MMM"}
var r100 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var r10 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var r1 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

func intToRoman(num int) string {
	return r1000[num/1000] +
		r100[(num/100)%10] +
		r10[(num/10)%10] +
		r1[num%10]
}
