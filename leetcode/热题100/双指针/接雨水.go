package 双指针

/*
对双指针的理解：

left从左向右遍历，right从右向左遍历；

则对left来说，leftMax一定准确，rightMax不一定准确，因为区间（left, right）的值还没有遍历，
但是left的rightMax一定 >= right的rightMax，所以只要leftMax < rightMax时，
我们不关心left的rightMax是多少了，因为它肯定比leftMax大，我们可以直接计算出left的存水量leftMax - nums[left];

对right来说，rightMax一定准确，leftMax不一定准确，因为区间（left, right）的值还没有遍历，
但是right的leftMax一定 >= left的leftMax，所以只要leftMax >= rightMax时，
我们不关心right的leftMax是多少了，因为它肯定比rightMax大，我们可以直接计算出right的存水量rightMax - nums[right];
*/

/*

 */

func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}
