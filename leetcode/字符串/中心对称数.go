package main

import "fmt"

/*
给定一个整数 n ，返回所有长度为 n 的 中心对称数 。你可以以 任何顺序 返回答案。

中心对称数 是一个数字在旋转了 180 度之后看起来依旧相同的数字（或者上下颠倒地看）。

输入：n = 2
输出：["11","69","88","96"]


输入：n = 1
输出：["0","1","8"]
*/

// 0 1 2 3 4 5 6 7 8 9

//初始化1
// 0 1 8

// 6 9

//111 916
//奇数，n=3
//1x1 8x8 9x6 6x9 10x01

//偶数，n=2
// 1''1 ...

func main() {
	s1 := rotate(3)
	fmt.Println(s1)
	s6 := rotate(15)
	//fmt.Println(rotate(6))
	for _, s := range s6 {
		if !isRotated(s) {
			fmt.Println(s, " is not rotated")
		}
	}
}

func rotate(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return res
	}

	if n%2 == 1 {
		res = append(res, "0", "1", "8")
	} else {
		res = append(res, "")
	}

	for i := n%2 + 2; i <= n; i += 2 {
		temp := make([]string, 0)
		for _, r := range res {
			temp = append(temp, "1"+r+"1")
			temp = append(temp, "8"+r+"8")
			temp = append(temp, "6"+r+"9")
			temp = append(temp, "9"+r+"6")
			if i != n {
				temp = append(temp, "0"+r+"0")
			}
		}

		res = temp
	}

	return res
}

func isRotated(s string) bool {
	pairs := map[string]string{"1": "1", "8": "8", "0": "0", "6": "9", "9": "6"}

	left, right := 0, len(s)-1
	for left < right {
		if pairs[string(s[left])] != string(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}
