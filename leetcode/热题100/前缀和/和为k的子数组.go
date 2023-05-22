package 前缀和

/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。



示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/subarray-sum-equals-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
维护一个map，存 当前元素前缀和：该前缀和出现次数的k：v

假设当前前缀和为preSum，若以前存在过前缀和为k的子数组，那么：preSum - k 必定在map中已经存在，将map[preSum-k]的出现次数累加
*/
func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	//默认preSum = 0已经出现过1次
	m[0] = 1
	preSum, res := 0, 0
	for _, n := range nums {
		preSum += n
		if v, ok := m[preSum-k]; ok {
			res += v
		}
		m[preSum] += 1
	}
	return res
}
