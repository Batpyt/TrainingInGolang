package 链表

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。



示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

 1 2 3 4 5

示例 2：

输入：head = [1], n = 1
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left, right := &ListNode{}, head
	left.Next = head
	dummy := left

	for n > 0 {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next
	return dummy.Next
}
