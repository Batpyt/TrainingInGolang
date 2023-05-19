package 双指针

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。



示例 1：



输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
双指针，left指向首，right指向尾。计算集水面积由更短的height和底部长度决定，
一开始的面积即为底部最长时的解。接下来，每次都将更短height那一边的指针往中心移动，并更新最大面积
*/

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := -1
	for left < right {
		res = max(res, area(height[left], height[right], right-left))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return res
}

func area(high1, high2, length int) int {
	if high1 > high2 {
		return high2 * length
	}
	return high1 * length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
