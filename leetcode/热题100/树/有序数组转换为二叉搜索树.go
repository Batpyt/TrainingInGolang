package main

func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

/*
每次都找到中间的节点当作根节点，递归一直到叶子结点
*/
func build(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	//如果除不开，定位到偏左的
	mid := (left + right) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}

	//递归构建左右子树
	root.Left = build(nums, left, mid-1)
	root.Right = build(nums, mid+1, right)
	return root
}
