package main

import "math"

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。



注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。


示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minWindow(s string, t string) string {
	exist := make(map[byte]int)
	mapt := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		mapt[t[i]]++
	}

	check := func() bool {
		for k, v := range mapt {
			if exist[k] < v {
				return false
			}
		}
		return true
	}

	length := math.MaxInt
	resL, resR := -1, -1

	for l, r := 0, 0; r < len(s); r++ {
		if mapt[s[r]] > 0 {
			exist[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < length {
				length = r - l + 1
				resL, resR = l, l+length
			}
			if _, ok := mapt[s[l]]; ok {
				exist[s[l]] -= 1
			}
			l++
		}
	}
	if resL == -1 {
		return ""
	}
	return s[resL:resR]
}
