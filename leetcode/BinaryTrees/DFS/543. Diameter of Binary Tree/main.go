package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		res = max(res, l+r)
		return max(l, r) + 1
	}

	_ = dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
