package 回溯

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。



示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
回溯法
*/
func permute(nums []int) [][]int {
	//维护三个变量，结果
	res := make([][]int, 0)
	//当前遍历的路径
	path := []int{}
	//当前遍历过的元素，已遍历过存入path的就标为true
	used := make([]bool, len(nums))

	var dfs func([]int, int)
	dfs = func(nums []int, cur int) {
		//cur表示当前path的长度，path长度==nums时就可以计入结果了
		if cur == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}

		//遍历nums的元素
		for i := 0; i < len(nums); i++ {
			//若当前位置的元素还没记录过，则走下面的逻辑
			if !used[i] {
				//元素计入path
				path = append(path, nums[i])
				//标记used已用过
				used[i] = true
				//下一层迭代
				dfs(nums, cur+1)
				//恢复现场
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}

	dfs(nums, 0)
	return res
}
