package lib

type Iterator struct {
	*ListNode
}

func (this *Iterator) HasNext() bool {
	// Returns true if the iteration has more elements.
	return this.ListNode != nil
}

func (this *Iterator) Next() int {
	// Returns the next element in the iteration.
	v := this.ListNode.Val
	this.ListNode = this.ListNode.Next
	return v
}
