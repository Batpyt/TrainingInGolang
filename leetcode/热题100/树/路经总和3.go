package main

func pathSum(root *TreeNode, targetSum int) int {
	//假设当前前缀和为preSum，若以前存在过前缀和为k的子数组，那么：preSum - k 必定在map中已经存在，将map[preSum-k]的出现次数累加
	preSum := make(map[int64]int)
	//默认0为1次
	preSum[0] = 1
	res := 0
	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, cur int64) {
		if node == nil {
			return
		}
		//将当前节点值加入当前前缀和
		cur += int64(node.Val)
		//
		res += preSum[cur-int64(targetSum)]
		preSum[cur]++
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		//恢复现场
		preSum[cur]--
		return
	}
	dfs(root, 0)
	return res
}
