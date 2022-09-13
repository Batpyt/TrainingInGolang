package main

import "strconv"

/*
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

示例 1 :

输入: 2736
输出: 7236
解释: 交换数字2和数字7。
示例 2 :

输入: 9973
输出: 9973
解释: 不需要交换。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-swap
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
贪心算法：每一位数字应该不小于所有排它后面的数字，否则找最大的且排最后面的数字与之交换

*/

func maximumSwap(num int) int {
	arr := []byte(strconv.Itoa(num))
	//maxindex代表最大数字的索引位置，index1，index2记录需要换位置的数字的索引
	maxIndex, index1, index2 := len(arr)-1, -1, -1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > arr[maxIndex] {
			//前面的数字大于后面的，将最大索引转移
			maxIndex = i
		} else if arr[i] < arr[maxIndex] {
			//后面的数字大于前面的，记录这两个需要换位置的索引
			index1, index2 = maxIndex, i
		}
	}
	if index1 < 0 {
		return num
	}
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
	res, _ := strconv.Atoi(string(arr))
	return res
}
