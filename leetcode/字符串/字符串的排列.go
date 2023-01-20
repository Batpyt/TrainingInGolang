package main

import (
	"sort"
)

//Mark：回溯

//func main() {
//	fmt.Println(permutation("abc"))
//}

/*
创建回溯函数，同时维护visit数组记录遍历到的元素，每次将没有遍历到的元素放入。
为了方便处理重复元素，先将原字符串按顺序排列，然后当碰见两相同元素时，必须按顺序放入。判断条件为两相同元素的前一个在不在visit中，若不在，跳过这一层循环
*/
func permutation(s string) []string {
	res := make([]string, 0)
	visit := make([]bool, len(s))
	temp := ""

	t := []byte(s)
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	var back func(int)
	back = func(i int) {
		if i == len(s) {
			res = append(res, temp)
		}
		for j, v := range visit {
			//当该元素已被遍历过 ｜｜ 碰见两相同元素时，必须按顺序放入。判断条件为两相同元素的前一个在不在visit中，若不在，跳过这一层循环
			if v || (j > 0 && !visit[j-1] && t[j-1] == t[j]) {
				continue
			}
			temp += string(t[j])
			visit[j] = true
			back(i + 1)
			visit[j] = false
			temp = temp[:len(temp)-1]
		}
	}
	back(0)
	return res
}
