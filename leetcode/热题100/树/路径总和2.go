package main

/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。



示例 1：


输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/path-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func pathSum2(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	path := []int{}

	//深度优先搜索，传递值为节点和targetSum的剩余值
	var dfs func(node *TreeNode, left int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}

		//计算剩余值
		left -= node.Val
		path = append(path, node.Val)

		//恢复现场
		defer func() {
			path = path[:len(path)-1]
		}()

		if node.Left == nil && node.Right == nil && left == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, targetSum)
	return res
}
