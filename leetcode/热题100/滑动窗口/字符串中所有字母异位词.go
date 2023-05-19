package main

import "fmt"

func main() {
	//fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}

/*
滑动窗口，转化成异位词问题。
*/
func findAnagrams(s string, p string) []int {
	left, right := 0, len(p)
	key := charsCount(p)
	res := make([]int, 0)

	for right <= len(s) {
		temp := s[left:right]
		tempKey := charsCount(temp)
		//字符出现次数数组相同的就是异位词
		if key == tempKey {
			res = append(res, left)
		}
		left++
		right++
	}
	return res
}

// 用26长度的数组存字符出现的次数
func charsCount(s string) [26]int {
	var key [26]int
	for _, ps := range s {
		key[ps-'a']++
	}
	return key
}
