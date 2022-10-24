package main

import "fmt"

/*
给你长度相等的两个字符串 s1 和 s2 。一次 字符串交换 操作的步骤如下：选出某个字符串中的两个下标（不必不同），并交换这两个下标所对应的字符。

如果对 其中一个字符串 执行 最多一次字符串交换 就可以使两个字符串相等，返回 true ；否则，返回 false 。



示例 1：

输入：s1 = "bank", s2 = "kanb"
输出：true
解释：例如，交换 s2 中的第一个和最后一个字符可以得到 "bank"
示例 2：

输入：s1 = "attack", s2 = "defend"
输出：false
解释：一次字符串交换无法使两个字符串相等
示例 3：

输入：s1 = "kelb", s2 = "kelb"
输出：true
解释：两个字符串已经相等，所以不需要进行字符串交换
示例 4：

输入：s1 = "abcd", s2 = "dcba"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/check-if-one-string-swap-can-make-strings-equal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var d1, d2 int
	var f1, f2 bool
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			fmt.Println(f1, f2)
			if !f1 && !f2 {
				d1 = i
				f1 = true
				continue
			}
			if f1 && !f2 {
				d2 = i
				f2 = true
				continue
			}
			if f1 && f2 {
				return false
			}
		}
	}

	fmt.Println(d1, d2)

	if s1[d1] == s2[d2] && s1[d2] == s2[d1] {
		return true
	}
	return false
}

func main() {
	fmt.Println(areAlmostEqual("bank", "kanb"))
}
