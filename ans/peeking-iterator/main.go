package main

import (
	"fmt"

	"github.com/kalokng/leetcode/lib"
)

func main() {
	l := lib.NewList(1, 2, 3)
	itr := Iterator(lib.Iterator{ListNode: l})
	pitr := PeekingIterator{Iterator: &itr}
	fmt.Println(pitr.hasNext())
	fmt.Println(pitr.next())
	fmt.Println(pitr.hasNext())
	fmt.Println(pitr.peek())
	fmt.Println(pitr.hasNext())
	fmt.Println(pitr.next())
	fmt.Println(pitr.hasNext())
	fmt.Println(pitr.peek())
	fmt.Println(pitr.peek())
	fmt.Println(pitr.hasNext())
	fmt.Println(pitr.next())
	fmt.Println(pitr.hasNext())
}

type Iterator lib.Iterator

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return (*lib.Iterator)(this).HasNext()
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return (*lib.Iterator)(this).Next()
}

type PeekingIterator struct {
	*Iterator
	done bool
	val  int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		Iterator: iter,
	}
}

func (this *PeekingIterator) hasNext() bool {
	if !this.done {
		return this.Iterator.hasNext()
	}
	return true
}

func (this *PeekingIterator) next() int {
	if !this.done {
		this.val = this.Iterator.next()
	}
	this.done = false
	return this.val
}

func (this *PeekingIterator) peek() int {
	if !this.done {
		this.val = this.Iterator.next()
		this.done = true
	}
	return this.val
}
