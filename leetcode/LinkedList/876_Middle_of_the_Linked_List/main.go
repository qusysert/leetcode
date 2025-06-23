package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	slow := head
	fast := head

	for {
		fast = fast.Next.Next
		if fast == nil || fast.Next == nil {
			return slow.Next
		}
		slow = slow.Next
	}
}
