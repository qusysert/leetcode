package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := NewQueue[*TreeNode]()
	q.Push(root)
	var sum int

	for q.Size() != 0 {
		levelLen := q.Size()
		sum = 0
		for range levelLen {
			node := q.Pop()
			sum += node.Val

			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}
	return sum
}

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Push(val T) {
	q.data = append(q.data, val)
}

func (q *Queue[T]) Pop() T {
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

//func deepestLeavesSum(root *TreeNode) int {
//	sum := 0
//	if root == nil {
//		return sum
//	}
//	return sum
//}
//
//func calc(root *TreeNode, sum *int) {
//	if root.Left == nil && root.Right == nil {
//		*sum += root.Val
//		return
//	}
//	calc(root.Left, sum)
//	calc(root.Right, sum)
//}
