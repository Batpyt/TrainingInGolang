package main

/*
给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。

请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。

下图中蓝色边和节点展示了操作后的结果：


请你返回结果链表的头指针。



示例 1：
输入：list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
输出：[0,1,2,1000000,1000001,1000002,5]

示例 2：
输入：list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
输出：[0,1,1000000,1000001,1000002,1000003,1000004,6]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/merge-in-between-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	root := &ListNode{}
	root.Next = list1
	left := &ListNode{}
	right := &ListNode{}
	for true {
		if a == 1 {
			left = list1
		}
		if b == 0 {
			right = list1.Next
			break
		}
		a--
		b--
		list1 = list1.Next
	}

	left.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = right
	return root.Next
}
