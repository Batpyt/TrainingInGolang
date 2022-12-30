package main

//func main() {
//	fmt.Println(exchange([]int{1, 2, 3, 4}))
//}

func exchange(nums []int) []int {
	res := make([]int, len(nums))
	left := 0
	right := len(res) - 1
	for _, n := range nums {
		if n%2 == 0 {
			res[right] = n
			right--
		} else {
			res[left] = n
			left++
		}
	}
	return res
}
