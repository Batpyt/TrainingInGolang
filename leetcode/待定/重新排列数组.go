package main

/*
给你一个数组 nums ，数组中有 2n 个元素，按 [x1,x2,...,xn,y1,y2,...,yn] 的格式排列。

请你将数组按 [x1,y1,x2,y2,...,xn,yn] 格式重新排列，返回重排后的数组。



示例 1：

输入：nums = [2,5,1,3,4,7], n = 3
输出：[2,3,5,4,1,7]
解释：由于 x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 ，所以答案为 [2,3,5,4,1,7]
示例 2：

输入：nums = [1,2,3,4,4,3,2,1], n = 4
输出：[1,4,2,3,3,2,4,1]
示例 3：

输入：nums = [1,1,2,2], n = 2
输出：[1,2,1,2]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/shuffle-the-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func shuffle(nums []int, n int) []int {
	first := 0
	res := make([]int, 0)
	for true {
		res = append(res, nums[first], nums[first+n])
		first++
		if first+n >= len(nums) {
			break
		}
	}
	return res
}

//func main() {
//	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
//}
