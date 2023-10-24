// https://leetcode.com/problems/binary-tree-preorder-traversal/description/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return []int{}
	}
	goThrough(root, &res)
	return res
}

func goThrough(root *TreeNode, res *[]int) {
	if root.Right != nil {
		goThrough(root.Right, res)
	}
	if root.Left != nil {
		goThrough(root.Left, res)
	}
	*res = append([]int{root.Val}, *res...)
}
