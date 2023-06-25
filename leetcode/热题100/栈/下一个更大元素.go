package main

/*
nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。

给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。

对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。

返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。



示例 1：

输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
输出：[-1,3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
示例 2：

输入：nums1 = [2,4], nums2 = [1,2,3,4].
输出：[3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/next-greater-element-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	//由于要求nums1在nums2中的结果，用一个map存nums2每个元素对应的结果
	resMap := make(map[int]int)

	//单调栈，栈中元素从栈顶到栈底永远是递增
	var stack []int

	//倒叙遍历nums2数组
	for i := len(nums2) - 1; i >= 0; i-- {
		temp := nums2[i]
		//若当前元素大于等于栈顶元素，为保证栈顶到栈底递增，将栈pop
		for len(stack) > 0 && temp >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		//此时若栈中还有元素，栈顶元素肯定>当前元素，将当前栈底记录为当前元素的结果
		if len(stack) > 0 {
			resMap[temp] = stack[len(stack)-1]
		} else { //栈中没有元素了的话，说明该元素右边没有大于它的，记录-1
			resMap[temp] = -1
		}
		stack = append(stack, temp)
	}

	res := make([]int, 0)
	for _, num := range nums1 {
		res = append(res, resMap[num])
	}
	return res
}
