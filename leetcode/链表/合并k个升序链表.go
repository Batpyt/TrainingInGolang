package main

import "fmt"

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}

	l3 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	lists := make([]*ListNode, 0)
	lists = append(lists, l1, l2, l3)
	res := mergeKLists(lists)

	print(res)
}

/*
两两合并
*/
func mergeKLists(lists []*ListNode) *ListNode {
	return mergeHelper(lists, 0, len(lists)-1)
}

func mergeHelper(listNodes []*ListNode, low, high int) *ListNode {
	if low == high {
		return listNodes[low]
	}

	if low > high {
		return nil
	}

	mid := (low + high) / 2
	return mergeTwoLists2(mergeHelper(listNodes, low, mid), mergeHelper(listNodes, mid+1, high))
}

/*
单纯的合并两个链表
*/
func mergeTwoLists2(headA *ListNode, headB *ListNode) *ListNode {
	//print(headA)
	//print(headB)
	if headA == nil {
		return headB
	}
	if headB == nil {
		return headA
	}

	dummy := &ListNode{}
	head := dummy
	for headA != nil && headB != nil {
		if headA.Val <= headB.Val {
			head.Next = headA
			headA = headA.Next
		} else {
			head.Next = headB
			headB = headB.Next
		}

		head = head.Next
	}

	if headA != nil {
		head.Next = headA
	}
	if headB != nil {
		head.Next = headB
	}
	//print(dummy.Next)
	return dummy.Next
}

func print(res *ListNode) {
	for res != nil {
		fmt.Print(res.Val, " -> ")
		res = res.Next
	}
	fmt.Println(" ")
}
