package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type n struct {
	Node *TreeNode
	Min  float64
	Max  float64
}

func isValidBST(root *TreeNode) bool {
	stack := NewStack[n]()
	stack.Push(n{root, math.Inf(-1), math.Inf(1)})

	for !stack.IsEmpty() {
		cur := stack.Pop()
		if !(float64(cur.Node.Val) > cur.Min && float64(cur.Node.Val) < cur.Max) {
			return false
		}
		if cur.Node.Left != nil {
			stack.Push(n{cur.Node.Left, float64(cur.Min), float64(cur.Node.Val)})
		}
		if cur.Node.Right != nil {
			stack.Push(n{cur.Node.Right, float64(cur.Node.Val), float64(cur.Max)})
		}
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
