package tree

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil {
		return root.Val == 1
	}
	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}
	return evaluateTree(root.Left) && evaluateTree(root.Right)
}
