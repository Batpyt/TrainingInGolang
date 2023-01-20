package main

import (
	"strings"
)

/*
一个句子是由一些单词与它们之间的单个空格组成，且句子的开头和结尾没有多余空格。比方说，"Hello World" ，"HELLO" ，"hello world hello world" 都是句子。每个单词都 只 包含大写和小写英文字母。

如果两个句子 sentence1 和 sentence2 ，可以通过往其中一个句子插入一个任意的句子（可以是空句子）而得到另一个句子，那么我们称这两个句子是 相似的 。比方说，sentence1 = "Hello my name is Jane" 且 sentence2 = "Hello Jane" ，我们可以往 sentence2 中 "Hello" 和 "Jane" 之间插入 "my name is" 得到 sentence1 。

给你两个句子 sentence1 和 sentence2 ，如果 sentence1 和 sentence2 是相似的，请你返回 true ，否则返回 false 。



示例 1：

输入：sentence1 = "My name is Haley", sentence2 = "My Haley"
输出：true
解释：可以往 sentence2 中 "My" 和 "Haley" 之间插入 "name is" ，得到 sentence1 。
示例 2：

输入：sentence1 = "of", sentence2 = "A lot of words"
输出：false
解释：没法往这两个句子中的一个句子只插入一个句子就得到另一个句子。
示例 3：

输入：sentence1 = "Eating right now", sentence2 = "Eating"
输出：true
解释：可以往 sentence2 的结尾插入 "right now" 得到 sentence1 。
示例 4：

输入：sentence1 = "Luky", sentence2 = "Lucccky"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sentence-similarity-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//func main() {
//	fmt.Println(areSentencesSimilar("My name is Haley", "My Haley"))
//	fmt.Println(areSentencesSimilar("A A AAa", "A AAa"))
//}

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	list1 := strings.Split(sentence1, " ")
	list2 := strings.Split(sentence2, " ")

	if len(list2) > len(list1) {
		list1, list2 = list2, list1
	}

	//fmt.Println(list1, list2)

	i := 0
	j1 := len(list1) - 1
	j2 := len(list2) - 1
	similar1, similar2 := 0, 0

	for i < len(list1) && i < len(list2) && list1[i] == list2[i] {
		i++
		similar1++
		if similar1 == len(list2) {
			return true
		}
	}

	for j1 >= 0 && j2 >= 0 && list1[j1] == list2[j2] {
		j1--
		j2--
		similar2++
		if similar2 == len(list2) || similar2+similar1 == len(list2) {
			return true
		}
	}
	return false
}
