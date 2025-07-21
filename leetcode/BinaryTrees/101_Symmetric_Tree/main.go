package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	stack := NewStack[*TreeNode]()
	stack.Push(root.Left)
	stack.Push(root.Right)

	for !stack.IsEmpty() {
		r := stack.Pop()
		l := stack.Pop()
		if (r == nil && l != nil) || (l == nil && r != nil) {
			return false
		}
		if r == nil && l == nil {
			continue
		}
		if r.Val != l.Val {
			return false
		}
		stack.Push(l.Left)
		stack.Push(r.Right)
		stack.Push(r.Left)
		stack.Push(l.Right)
	}
	return true
}

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() T {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
