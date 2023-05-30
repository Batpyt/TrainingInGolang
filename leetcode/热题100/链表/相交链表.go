package 链表

/*
	4,1,8,4,5

5,6,1,8,4,5
查看以上例子，若两链表会相交，且相交后的长度为c，各自相交前的长度为a、b，则：
a+c+b == b+c+a，当两个链表都移动了相同距离后，在交点相遇
若不相交，长度分别为a，b，则：
a+b = b+a，相同距离后都移动到了尾节点也就是nil
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB

	for true {
		//两种情况，相遇或都移动到了尾
		if a == b {
			break
		}
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}
