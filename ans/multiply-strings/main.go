package main

import "fmt"

func main() {
	fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("123", "456"))
	fmt.Println(multiply("0", "0"))
	fmt.Println(multiply("999", "0"))
	fmt.Println(multiply("999", "999"))
	fmt.Println(multiply("0", "999"))
	fmt.Println(multiply("999999999999", "999999999"))
	fmt.Println(multiply("0002", "0030"))
}

func multiply(num1 string, num2 string) string {
	var d []byte
	var n1 []byte
	var n2 []byte
	for i := len(num1) - 1; i >= 0; i-- {
		n1 = append(n1, num1[i]-'0')
	}
	for i := len(num2) - 1; i >= 0; i-- {
		n2 = append(n2, num2[i]-'0')
	}

	for i, v := range n1 {
		for j, v2 := range n2 {
			if len(d) <= i+j {
				d = append(d, 0)
			}
			r := d[i+j] + v*v2
			if r >= 10 {
				if len(d) <= i+j+1 {
					d = append(d, 0)
				}
				d[i+j+1] = d[i+j+1] + r/10
			}
			d[i+j] = r % 10
		}
	}
	for i := len(d) - 1; i > 0; i-- {
		if d[i] != 0 {
			break
		}
		d = d[:i]
	}
	i := 0
	j := len(d) - 1
	for i < j {
		d[i], d[j] = d[j], d[i]
		i++
		j--
	}
	for i := range d {
		d[i] += '0'
	}
	return string(d)
}
