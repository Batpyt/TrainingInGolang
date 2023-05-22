package main

import "fmt"

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。



示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	/*
		单调队列
		单调队列存储nums元素的下标，规则如下：
		下标的顺序按照从小到大存在队列中，但是下标对应的nums[]元素要从大到小。
		也就是当push一个新值i进来时，要将所有nums[j] <= nums[i]的 j 都移出队列，然后将i append进入q

		也就是：若 i < j 那么 nums[i] > nums[j]
		因此，q[0]永远是nums中最靠左的最大值的下标，可以看作一个最大堆的堆顶
	*/
	q := make([]int, 0)
	push := func(i int) {
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	//初始化第一个窗口，此时的q[0]就指向了该窗口内最大值的下标
	for i := 0; i < k; i++ {
		push(i)
	}
	res := make([]int, 0)
	res = append(res, nums[q[0]])

	/*
		遍历数组。q[0]永远记录的是最新窗口内的最大元素下标.
		当有新的值 nums[i] 进入窗口时，将所有 nums[j] <= nums[i] 的j移出q
		将被移出窗口的下标移出q（如果存在）
	*/
	for i := k; i < len(nums); i++ {
		push(i)
		//将已经不在窗口范围内的下标移出队列
		for q[0] <= i-k {
			q = q[1:]
		}
		res = append(res, nums[q[0]])
	}
	return res
}
