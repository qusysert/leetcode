// https://leetcode.com/problems/design-linked-list/
package main

func main() {
	ll := Constructor()

	ll.AddAtHead(1)
	ll.AddAtTail(3)
	ll.AddAtIndex(1, 2)
	ll.Get(1)
	ll.DeleteAtIndex(1)
	ll.Get(1)
}

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}
	if this.size == 0 {
		this.head = node
		this.tail = node
		this.size++
		return
	}
	node.Next = this.head
	this.head = node
	this.size++

}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}
	if this.size == 0 {
		this.head = node
		this.tail = node
		this.size++
		return
	}
	this.tail.Next = node
	this.tail = node
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}
	this.size++
	node := &Node{Val: val}
	prev := this.head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}
	node.Next = prev.Next
	prev.Next = node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	if index == 0 {
		this.head = this.head.Next
		if this.head == nil {
			this.tail = nil
		}
		this.size--
		return
	}

	prev := this.head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}
	this.size--
	if prev.Next == this.tail {
		this.tail = prev
		prev.Next = nil
		return
	}
	prev.Next = prev.Next.Next
}
