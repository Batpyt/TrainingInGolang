package main

/*
	存成 num：index的字典，为保证不会自己与自己匹配，只有遍历过自己以后才将自己存入字典
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, index}
		}
		m[num] = index
	}
	return []int{}
}
