package main

import "math"

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt, math.MaxInt)
}
func validate(node *TreeNode, lower int, upper int) bool {
	if node == nil {
		return true
	}
	if (lower < node.Val) && (upper > node.Val) {
		return validate(node.Left, lower, node.Val) && validate(node.Right, node.Val, upper)
	} else {
		return false
	}

}
