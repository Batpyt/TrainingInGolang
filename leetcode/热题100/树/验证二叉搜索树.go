package main

import "math"

/*
		10
	  /    \
	 5		15
		   /  \
		  6   20

以这个树为例
15的节点，lower为10，意味着15以后的节点都必须大于10，upper还没变
5的节点，upper为10，意味着5以后的节点都必须小雨10，lower还没变
由此可知，左子树都关心upper,upper随着递归更改；
右子树都关心lower，lower随着递归更改
*/
func isValidBST(root *TreeNode) bool {
	return test(root, math.MinInt64, math.MaxInt64)
}

func test(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return test(root.Left, lower, root.Val) && test(root.Right, root.Val, upper)
}
