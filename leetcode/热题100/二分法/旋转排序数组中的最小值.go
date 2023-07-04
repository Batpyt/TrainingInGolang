package 二分法

/*
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2

		//当nums[pivot] < nums[high]，说明nums[pivot]是最小值右侧的元素，右指针前移
		if nums[pivot] < nums[high] {
			//为什么此处不+1？因为如果pivot位置的元素就是最小值本身，这样就会错过
			high = pivot
		} else {
			//当nums[pivot] > nums[high]，说明nums[pivot]是最小值左侧的元素，左指针后移
			//为什么这里+1？因为nums[pivot] > nums[high]，所以pivot处不可能是最小值
			low = pivot + 1
		}
	}
	/*
		为什么不是low <= high 跳出
		low<high，假如最后循环到{*,10,1,*}的这种情况时，nums[low]=10,nums[high]=1,nums[mid]=10,low=mid+1,
		直接可以跳出循环了,所以low<high,此时low指向的就是最小值的下标;
		如果low<=high的话，low=high，还会再不必要的循环一次，此时最后一次循环的时候会发生low==high==mid，
		则nums[mid]==nums[high]，则会走一次else语句，则low=mid+1,此时low指向的是最小值的下一个下标，
		则需要return[low-1]
	*/
	return nums[low]
}
