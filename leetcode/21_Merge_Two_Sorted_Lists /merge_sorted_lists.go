// https://leetcode.com/problems/merge-two-sorted-lists/description/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	f := list1
	s := list2

	head := ListNode{}
	cur := &head
	for f != nil && s != nil {
		if f.Val < s.Val {
			cur.Next = f
			f = f.Next
		} else {
			cur.Next = s
			s = s.Next
		}
		cur = cur.Next
	}
	if f == nil {
		cur.Next = s
	}
	if s == nil {
		cur.Next = f
	}
	return head.Next
}
