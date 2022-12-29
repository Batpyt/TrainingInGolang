package main

/*
给定一个数组 nums ，将其划分为两个连续子数组 left 和 right， 使得：

left 中的每个元素都小于或等于 right 中的每个元素。
left 和 right 都是非空的。
left 的长度要尽可能小。
在完成这样的分组后返回 left 的 长度 。

用例可以保证存在这样的划分方法。



示例 1：

输入：nums = [5,0,3,8,6]
输出：3
解释：left = [5,0,3]，right = [8,6]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/partition-array-into-disjoint-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func partitionDisjoint(nums []int) int {
	n := len(nums)
	//维护三个变量，左边数组最大值，左边数组边界，当前最大值
	leftMax, leftPos, curMax := nums[0], 0, nums[0]
	for i := 1; i < n-1; i++ {
		//每往后一个都计算出当前最大值
		curMax = max(curMax, nums[i])
		//当当前遍历到的数小于左边数组最大值时，要把当前的数归于左边数组
		if nums[i] < leftMax {
			leftMax = curMax
			leftPos = i
		}
	}
	return leftPos + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
