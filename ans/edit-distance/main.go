package main

import "fmt"

type dmap struct {
	diff   []int
	w      int
	w1, w2 string
}

func newDmap(w1, w2 string) *dmap {
	d := &dmap{
		w:  len(w2) + 1,
		w1: w1,
		w2: w2,
	}
	d.diff = make([]int, (len(w1)+1)*(len(w2)+1))
	for i := range d.diff {
		d.diff[i] = -1
	}
	return d
}

func (d *dmap) set(i, j, v int) {
	d.diff[i*d.w+j] = v
}
func (d *dmap) val(i, j int) (int, bool) {
	v := d.diff[i*d.w+j]
	return v, v >= 0
}

func minDistance(word1 string, word2 string) int {
	d := newDmap(word1, word2)
	for i := 0; i <= len(word1); i++ {
		d.set(i, 0, i)
	}
	for i := 0; i <= len(word2); i++ {
		d.set(0, i, i)
	}
	return dist(d, len(word1), len(word2))
}

func dist(d *dmap, w1, w2 int) (diff int) {
	if v, ok := d.val(w1, w2); ok {
		return v
	}
	defer func() {
		d.set(w1, w2, diff)
	}()

	if d.w1[w1-1] == d.w2[w2-1] {
		return dist(d, w1-1, w2-1)
	}
	min := dist(d, w1-1, w2)
	if v := dist(d, w1, w2-1); v < min {
		min = v
	}
	if v := dist(d, w1-1, w2-1); v < min {
		min = v
	}
	return min + 1
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}
