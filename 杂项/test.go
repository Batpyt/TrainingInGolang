package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[0] = 0

	nums := []int{1, 2, 3}
	nums2 := make([]int, 0)
	nums2 = append(nums2, 1)
	change(m, nums, &nums2)
	fmt.Println(m, nums, nums2)
	//map[0:1 1:1] [2 3 4] [1 3]
	/*
		只有当方法的传参是数组的指针时数组的append才会生效
		但是++可以生效
	*/
}

func change(m map[int]int, nums []int, nums2 *[]int) {
	m[0] = 1
	m[1] = 1

	for i := range nums {
		nums[i]++
	}
	nums = append(nums, 4)
	*nums2 = append(*nums2, 3)
}
