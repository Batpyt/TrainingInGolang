package main

/*
 1-2-3-4-5
*/
/*
https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
*/
func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	_ = removeNthFromEnd(l, 2)
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := &ListNode{}
	quick := head
	slow.Next = head
	dummy := slow
	for n > 1 {
		quick = quick.Next
		n--
	}

	for quick.Next != nil {
		quick = quick.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
