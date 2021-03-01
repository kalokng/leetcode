package main

import "fmt"

func main() {
	s := Constructor()
	s.Push(5)
	s.Push(7)
	s.Push(5)
	s.Push(7)
	s.Push(4)
	s.Push(5)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

type FreqStack struct {
	lvl  [][]int
	freq map[int]int
}

func Constructor() FreqStack {
	return FreqStack{
		freq: make(map[int]int),
	}
}

func (this *FreqStack) Push(x int) {
	l := this.freq[x]
	for len(this.lvl) <= l {
		this.lvl = append(this.lvl, nil)
	}
	this.lvl[l] = append(this.lvl[l], x)
	this.freq[x] = l + 1
}

func (this *FreqStack) Pop() int {
	s := this.lvl[len(this.lvl)-1]
	x := s[len(s)-1]
	if len(s) == 1 {
		this.lvl = this.lvl[:len(this.lvl)-1]
	} else {
		s = s[:len(s)-1]
		this.lvl[len(this.lvl)-1] = s
	}
	this.freq[x]--
	return x
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */
