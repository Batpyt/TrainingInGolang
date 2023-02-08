package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{3, 4, 7, 2, -3, 1, 4, 2, 7}, 7))
}

/*
前缀和+map
用map记录前缀和和它出现的次数
当前前缀和 preSum
若当前preSum - k 在map中存在，则存在的出现的次数则为改元素位置为尾的符合条件的子数组数量，累加。

假设一个元素的前缀和presum1，当这之后的子数组和为k时，k位置的前缀和为：presum1 + k
因此，presum2 - k = presum1，所以当presum2 - 开出现过时，说明当前存在和为k的子数组
*/
func subarraySum(nums []int, k int) int {
	var preSum, res int
	exist := make(map[int]int)
	exist[0] = 1
	for _, n := range nums {
		preSum += n
		if _, ok := exist[preSum-k]; ok {
			res += exist[preSum-k]
		}
		fmt.Println(n, preSum, preSum-k, exist, res)
		exist[preSum] += 1
	}
	return res
}
