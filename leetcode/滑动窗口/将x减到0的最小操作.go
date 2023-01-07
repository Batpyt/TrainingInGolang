package 滑动窗口

/*
给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数组以供接下来的操作使用。

如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。



示例 1：

输入：nums = [1,1,4,2,3], x = 5
输出：2
解释：最佳解决方案是移除后两个元素，将 x 减到 0 。
示例 2：

输入：nums = [5,6,7,8,9], x = 4
输出：-1
示例 3：

输入：nums = [3,2,20,1,1,3], x = 10
输出：5
解释：最佳解决方案是移除后三个元素和前两个元素（总共 5 次操作），将 x 减到 0 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
滑动窗口。维护两个指针，left为左边减去的元素下标，right为右边的。lsum，rsum分别为左右两边减掉的和。
*/
func minOperations(nums []int, x int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum < x {
		return -1
	}

	//起初设置left为-1，right为0，意为左边不减，右边减去全部。
	lsum := 0
	left := -1
	rsum := sum
	right := 0
	res := len(nums) + 1
	for left < len(nums) {
		if left >= 0 {
			lsum += nums[left]
		}

		//当减去的值大于x时，需要把右指针往右，使减去的值减小
		for right < len(nums) && lsum+rsum > x {
			rsum -= nums[right]
			right++
		}
		if lsum+rsum == x {
			res = min(res, left+1+len(nums)-right)
		}
		left++
	}
	if res > len(nums) {
		return -1
	}
	return res
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
