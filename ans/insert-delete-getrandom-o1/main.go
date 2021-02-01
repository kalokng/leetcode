package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	idx map[int]int
	set []int
}

const size = 100000

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		idx: make(map[int]int),
		set: make([]int, 0, size),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idx[val]; ok {
		return false
	}
	i := len(this.set)
	this.idx[val] = i
	this.set = append(this.set, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.idx[val]
	if !ok {
		return false
	}
	l := len(this.set) - 1
	if l >= 0 && l != i {
		this.set[i] = this.set[l]
		this.idx[this.set[i]] = i
	}
	this.set = this.set[:l]
	delete(this.idx, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	l := len(this.set)
	return this.set[rand.Intn(l)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	o := Constructor()
	fmt.Println(o.Insert(1))
	fmt.Println(o)
	fmt.Println(o.Insert(2))
	fmt.Println(o)
	fmt.Println(o.Insert(3))
	fmt.Println(o)
	fmt.Println(o.GetRandom())
	fmt.Println(o.GetRandom())
	fmt.Println(o.GetRandom())
	fmt.Println(o.GetRandom())
	fmt.Println(o.GetRandom())
	fmt.Println(o.GetRandom())
	fmt.Println(o.GetRandom())
	fmt.Println(o.GetRandom())
	fmt.Println(o.GetRandom())
	fmt.Println(o.GetRandom())
	fmt.Println(o.GetRandom())
	fmt.Println(o.GetRandom())
	fmt.Println(o)
	fmt.Println(o.Remove(1))
	fmt.Println(o)
	fmt.Println(o.Insert(2))
	fmt.Println(o)
	fmt.Println(o.GetRandom())
}
