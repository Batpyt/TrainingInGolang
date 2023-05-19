package main

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。



示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
由于字母异位词的字符出现次数相等，因此将字符串的每个字符出现次数当作key，存入字典中。
*/
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		temp := [26]int{}
		for _, s := range str {
			temp[s-'a']++
		}
		m[temp] = append(m[temp], str)
	}

	res := make([][]string, 0, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
