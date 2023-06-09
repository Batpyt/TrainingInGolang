package main

/*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。



示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	//维护队列存放每一层的节点。一开始是根节点
	q := []*TreeNode{root}
	//i是层数，当队列中还有节点时就继续循环
	for i := 0; len(q) > 0; i++ {
		//不这么写直接append数组的话，会报oom
		ret = append(ret, []int{})
		//存储该层节点子节点的数组
		var p []*TreeNode
		//遍历该层节点，记录左右子节点，记录val
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		//将该层左右子节点记录至队列
		q = p
	}
	return ret
}
