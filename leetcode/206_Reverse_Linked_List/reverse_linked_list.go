// https://leetcode.com/problems/reverse-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := ListNode{head.Val, nil}
	reversed := &h
	for {
		prev := ListNode{head.Next.Val, nil}
		prev.Next = reversed
		reversed = &prev
		if head.Next.Next != nil {
			head = head.Next
		} else {
			return reversed
		}
	}
}
