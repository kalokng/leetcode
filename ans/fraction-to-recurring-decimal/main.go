package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	var neg bool
	if numerator < 0 {
		numerator = -numerator
		neg = !neg
	}
	if denominator < 0 {
		denominator = -denominator
		neg = !neg
	}
	first := numerator / denominator
	numerator = numerator % denominator
	if numerator == 0 {
		if neg {
			first = -first
		}
		return strconv.Itoa(first)
	}
	var sign string
	if neg {
		sign = "-"
	}
	var digit []byte
	m := map[int]int{numerator: 0}
	i := 0
	for numerator != 0 {
		numerator *= 10
		digit = append(digit, '0'+byte(numerator/denominator))
		numerator = numerator % denominator
		if j, ok := m[numerator]; ok {
			return sign + fmt.Sprintf("%d.%s(%s)", first, digit[:j], digit[j:])
		}
		i++
		m[numerator] = i
	}
	return sign + fmt.Sprintf("%d.%s", first, digit)
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(1, 7))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(4, 333))
	fmt.Println(fractionToDecimal(1, 5))
	fmt.Println(fractionToDecimal(1, 8))
	fmt.Println(fractionToDecimal(1, 1024))
	fmt.Println(fractionToDecimal(7, 990))
	fmt.Println(fractionToDecimal(2, 70))
	fmt.Println(fractionToDecimal(22, 7))
	fmt.Println(fractionToDecimal(355, 113))
	fmt.Println(fractionToDecimal(355, 113000))
	fmt.Println(fractionToDecimal(-50, 8))
	fmt.Println(fractionToDecimal(50, -8))
	fmt.Println(fractionToDecimal(-50, -8))
	fmt.Println(fractionToDecimal(7, -12))
	fmt.Println(fractionToDecimal(-12, 1))
	fmt.Println(fractionToDecimal(-2147483648, 1))
	fmt.Println(fractionToDecimal(0, -5))
	fmt.Println(fractionToDecimal(-2147483648, -1))
	fmt.Println(fractionToDecimal(0, 1))
	fmt.Println(fractionToDecimal(0, -1))
}
