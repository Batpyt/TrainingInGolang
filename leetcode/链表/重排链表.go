package main

/*
输入：head = [1,2,3,4]
输出：[1,4,2,3]
示例 2：



输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]
*/

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	nodes := make([]*ListNode, 0)
	for n := head; n != nil; n = n.Next {
		nodes = append(nodes, n)
	}

	left := 0
	right := len(nodes) - 1
	for left < right {
		nodes[left].Next = nodes[right]
		left++
		if left >= right {
			break
		}
		nodes[right].Next = nodes[left]
		right--
	}
	nodes[left].Next = nil
}
