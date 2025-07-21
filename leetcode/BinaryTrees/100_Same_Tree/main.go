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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack := NewStack[*TreeNode]()

	stack.Push(q)
	stack.Push(p)

	for !stack.IsEmpty() {
		p = stack.Pop()
		q = stack.Pop()

		if p == nil && q == nil {
			continue
		}
		if (q == nil && p != nil) || (p == nil && q != nil) {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		stack.Push(q.Left)
		stack.Push(p.Left)
		stack.Push(q.Right)
		stack.Push(p.Right)
	}
	return true
}

type Stack[T any] struct {
	stack []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	s.stack = append(s.stack, val)
}

func (s *Stack[T]) Pop() T {
	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return value
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}
