package 链表

/*
因此，当发现slow 与fast 相遇时，我们再额外使用一个指针ptr。起始，它指向链表头部；随后，它和slow 每次向后移动一个位置。最终，它们会在入环点相遇。

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
