package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
给定链表头结点 head，该链表上的每个结点都有一个 唯一的整型值 。同时给定列表 nums，该列表是上述链表中整型值的一个子集。

返回列表 nums 中组件的个数，这里对组件的定义为：链表中一段最长连续结点的值（该值必须在列表 nums 中）构成的集合。



示例 1：



输入: head = [0,1,2,3], nums = [0,1,3]
输出: 2
解释: 链表中,0 和 1 是相连接的，且 nums 中不包含 2，所以 [0, 1] 是 nums 的一个组件，同理 [3] 也是一个组件，故返回 2。
示例 2：



输入: head = [0,1,2,3,4], nums = [0,3,1,4]
输出: 2
解释: 链表中，0 和 1 是相连接的，3 和 4 是相连接的，所以 [0, 1] 和 [3, 4] 是两个组件，故返回 2。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/linked-list-components
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
就是找nums里的数能在链表中组成多少段
*/
func numComponents(head *ListNode, nums []int) int {
	set := make(map[int]bool)
	for _, n := range nums {
		set[n] = true
	}

	res := 0
	for inSet := false; head != nil; head = head.Next {
		if _, ok := set[head.Val]; !ok { //结束一段
			inSet = false
			continue
		}
		if !inSet { //开始新的一段
			inSet = true
			res++
		}
	}
	return res
}
