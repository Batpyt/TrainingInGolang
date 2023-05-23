package main

/*
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/reverse-nodes-in-k-group
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
https://leetcode.cn/problems/reverse-nodes-in-k-group/solution/tu-jie-kge-yi-zu-fan-zhuan-lian-biao-by-user7208t/
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	/*
		通过三个变量完成两个链表节点（prev，head）的反转
		prev -> head -> next
		prev <- head    next
		[]   <- prev    head
	*/
	reverse := func(head *ListNode) *ListNode {
		pre := &ListNode{}
		cur := head
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		//此时的prev变成了首节点
		return pre
	}

	dummy := &ListNode{
		Next: head,
	}

	/*
		需要两个初始变量，pre为需要翻转片段的前一个节点，end为需要翻转的片段的尾节点
		在翻转过程中需要start表示需要翻转的片段的首节点，next为需要翻转片段的后一个节点
		pre -> start -> ... -> end -> next
	*/
	pre := dummy
	end := dummy

	for end.Next != nil {
		//确定需要翻转的片段的尾部
		for i := 0; i < k; i++ {
			end = end.Next
			//若已到整个链表的尾，直接返回结果
			if end == nil {
				return dummy.Next
			}
		}

		//根据pre和end确定start和next的位置
		start := pre.Next
		next := end.Next
		//切断end和next的连接，迎合reverse方法的判断条件，将end暂且视为一个尾
		end.Next = nil
		//reverse方程返回的是翻转后的首部，pre指向该节点
		pre.Next = reverse(start)
		//此时start节点被翻转到了尾部，将其指向原本的next
		start.Next = next
		//将pre和end移动到start的位置，作为下一次循环的起始位置
		pre = start
		end = pre
	}
	return dummy.Next
}
