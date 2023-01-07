package tree

func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	level := make([]*TreeNode, 0)
	level = append(level, root)

	for len(level) > 0 {
		nodes := level
		level = nil
		for _, node := range nodes {
			res = append(res, node.Val)
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
	}
	return res
}

func levelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	level := make([]*TreeNode, 0)
	level = append(level, root)

	for len(level) > 0 {
		nodes := level
		level = nil
		temp := make([]int, 0)
		for _, node := range nodes {
			temp = append(temp, node.Val)
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

func levelOrder3(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	level := make([]*TreeNode, 0)
	level = append(level, root)

	index := 0
	for len(level) > 0 {
		nodes := level
		level = nil
		temp := make([]int, 0)
		for _, node := range nodes {
			temp = append(temp, node.Val)
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if index%2 == 1 {
			for i, n := 0, len(temp); i < n/2; i++ {
				temp[i], temp[n-1-i] = temp[n-1-i], temp[i]
			}
		}
		index++
		res = append(res, temp)
	}
	return res
}
