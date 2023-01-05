package main

import (
	"strconv"
	"strings"
)

//func main() {
//	fmt.Println(areNumbersAscending("1 2 uioioi 4 5 ioioopp 6"))
//	fmt.Println(areNumbersAscending("1 4 uioioi 4 5 ioioopp 6"))
//}

func areNumbersAscending(s string) bool {
	list := strings.Split(s, " ")
	nums := make([]int64, 0)
	for _, l := range list {
		if num, err := strconv.ParseInt(l, 10, 64); err == nil {
			nums = append(nums, num)
		}
	}
	prev := int64(-1)
	for _, num := range nums {
		if prev >= num {
			return false
		}
		prev = num
	}
	return true
}
