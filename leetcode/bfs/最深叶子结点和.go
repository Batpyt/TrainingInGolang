package bfs

/*
给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。


输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
输出：15

链接：https://leetcode.cn/problems/deepest-leaves-sum
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	q := []*TreeNode{root}
	var sum int
	for len(q) > 0 {
		sum = 0
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			sum = sum + node.Val

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return sum
}