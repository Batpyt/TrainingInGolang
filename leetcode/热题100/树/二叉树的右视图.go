package main

/*
层序遍历，每层的最右边的元素
*/
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		var p []*TreeNode
		temp := make([]int, 0)
		for j := 0; j < len(q); j++ {
			node := q[j]
			temp = append(temp, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
		//记录每层最右边元素
		if len(temp) > 0 {
			res = append(res, temp[len(temp)-1])
		}
	}
	return res
}
