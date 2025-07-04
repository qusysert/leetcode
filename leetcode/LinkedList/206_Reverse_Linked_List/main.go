package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head
	var nxt *ListNode

	for cur != nil {
		nxt = cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	return prev
}
