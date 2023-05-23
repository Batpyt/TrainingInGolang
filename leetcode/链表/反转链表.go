package main

/*
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func reverseList(head *ListNode) *ListNode {
	//刚开始时，需要假节点来完成首节点的翻转
	var prev *ListNode
	/*
		通过三个变量完成两个链表节点（prev，head）的反转
		prev -> head -> next
		prev <- head    next
		[]   <- prev    head
	*/
	for head != nil {
		//先将head的原本next节点保留
		next := head.Next
		//翻转head的next指针，指向prev
		head.Next = prev
		//prev和head往后平移
		prev = head
		head = next
	}
	return prev
}
