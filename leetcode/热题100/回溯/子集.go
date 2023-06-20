package 回溯

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。



示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := []int{}

	var dfs func(int)
	dfs = func(cur int) {
		//cur代表遍历到了第几个元素，当遍历到最后一个元素时，记录结果并返回
		if cur == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		path = append(path, nums[cur])
		//考虑当前位置
		dfs(cur + 1)
		//不考虑当前位置，将最后一位删除
		path = path[:len(path)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return res
}
