package 链表

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	for ; k > 0; k-- {
		fast = fast.Next
	}

	for fast != nil {
		head = head.Next
		fast = fast.Next
	}
	return head
}
