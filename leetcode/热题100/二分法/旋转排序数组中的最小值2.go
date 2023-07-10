package 二分法

/*
与原题相比，这次要考虑重复元素
*/

func findMin2(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else if nums[pivot] > nums[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return nums[low]
}
