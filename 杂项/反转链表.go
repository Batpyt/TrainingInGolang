package main

import (
	"TrainingInGolang2/leetcode/热题100/链表"
)

//
//func main() {
//	head := &链表.ListNode{
//		Val:  1,
//		Next: &链表.ListNode{
//			Val:  2,
//			Next: &链表.ListNode{
//				Val:  3,
//				Next: &链表.ListNode{
//					Val:  4,
//					Next: nil,
//				},
//			},
//		},
//	}
//
//	res := reverseList(head)
//
//	for res != nil {
//		fmt.Print(res.Val, " -> ")
//		res = res.Next
//	}
//}

/*
 pre -> 1 -> 2 -> 3 -> 4
 pre <- cur

*/
func reverseList(head *链表.ListNode) *链表.ListNode {
	if head == nil {
		return nil
	}

	var pre *链表.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
