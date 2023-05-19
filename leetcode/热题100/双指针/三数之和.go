package 双指针

import (
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
先排序，然后将问题转化为双指针问题
-a = b + c
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	//第一个指针，定位a
	for first := 0; first < len(nums); first++ {
		//若与上一个相同则跳过
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		//第三个指针，定位c，且从后往前移动
		third := len(nums) - 1
		target := -1 * nums[first]

		//第二个指针，定位b，从a的后一个开始向后移动
		for second := first + 1; second < len(nums); second++ {
			//相同则跳过
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			//若 -a < b + c，将c的指针继续向前移动
			for second < third && nums[second]+nums[third] > target {
				third--
			}

			//若二三指针相遇，则未找到结果，跳出循环
			if second == third {
				break
			}

			//记录符合条件的结果
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}

//func main() {
//	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
//}

// -a = b + c
func th(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, 0)

	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := len(nums) - 1
		target := -1 * nums[first]

		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for target < nums[second]+nums[third] && second < third {
				third--
			}

			if second == third {
				break
			}

			if target == nums[second]+nums[third] {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}
