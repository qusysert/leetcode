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
type leaf struct {
	Node *TreeNode
	Sum  int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := NewStack[leaf]()
	stack.Push(leaf{root, root.Val})

	for !stack.IsEmpty() {
		cur := stack.Pop()
		if cur.Sum == targetSum && cur.Node.Left == nil && cur.Node.Right == nil {
			return true
		}

		if cur.Node.Left != nil {
			stack.Push(leaf{cur.Node.Left, cur.Sum + cur.Node.Left.Val})
		}
		if cur.Node.Right != nil {
			stack.Push(leaf{cur.Node.Right, cur.Sum + cur.Node.Right.Val})
		}
	}
	return false
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
