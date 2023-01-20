package tree

/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。



示例 1：



输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//Mark:回溯
func pathSum(root *TreeNode, target int) [][]int {
	res := make([][]int, 0)

	//记录遍历到的节点的数组
	path := make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		//减去该节点的值
		sum -= node.Val
		//将该节点记录至数组
		path = append(path, node.Val)
		//回溯
		defer func() {
			path = path[:len(path)-1]
		}()

		//遍历到了根节点且遍历过的值和为目标值，记录并返回
		if node.Left == nil && node.Right == nil && sum == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, target)
	return res
}
