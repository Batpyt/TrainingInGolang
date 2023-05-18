package main

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
用字典将数组去重，为追求最小时间复杂度，用以下逻辑进行求解
作为最长的连续序列，序列首x，在原数组中一定不存在x-1，因此，在原数组中遍历，当遍历到不存在x-1的元素时，
开始遍历x++，找到以x为首序列的最长序列。不断寻找并记录更长的序列。
*/

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)

	for _, n := range nums {
		m[n] = true
	}

	maxCount := 0
	for _, num := range nums {
		if !m[num-1] {
			currentNum := num
			currentCount := 1
			for m[currentNum+1] == true {
				currentNum++
				currentCount++
			}
			maxCount = max(maxCount, currentCount)
		}
	}
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
