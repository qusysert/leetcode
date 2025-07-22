package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	q := NewQueue[*TreeNode]()
	q.Push(root)

	for !q.IsEmpty() {
		levelSize := q.Size()
		cur := make([]int, 0, levelSize)
		for range levelSize {
			node := q.Pop()
			cur = append(cur, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}

		res = append(res, getMaxInt(cur))
	}
	return res
}

func getMaxInt(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	max := slice[0]
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Push(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Pop() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}
