package main

//func main() {
//	fmt.Println(twoOutOfThree([]int{1, 1, 3, 2}, []int{3, 2}, []int{3}))
//}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	new1 := quchong(nums1)
	new2 := quchong(nums2)
	new3 := quchong(nums3)
	m := make(map[int]bool)
	nums := new1
	nums = append(nums, new2...)
	nums = append(nums, new3...)
	res := make([]int, 0)
	exist := make(map[int]bool)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			if _, ok2 := exist[n]; !ok2 {
				res = append(res, n)
				exist[n] = true
			}
			continue
		}
		m[n] = true
	}
	return res
}

func quchong(nums []int) []int {
	m := make(map[int]bool)
	res := make([]int, 0)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			continue
		}
		res = append(res, n)
		m[n] = true
	}
	return res
}
