package main

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

type leaf struct {
	node  *TreeNode
	depth int
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	stack := NewStack[leaf]()
	stack.Push(leaf{root, 1})
	for !stack.IsEmpty() {
		cur := stack.Pop()
		res = max(res, cur.depth)
		if cur.node.Left != nil {
			stack.Push(leaf{cur.node.Left, cur.depth + 1})
		}
		if cur.node.Right != nil {
			stack.Push(leaf{cur.node.Right, cur.depth + 1})
		}
	}
	return res
}
