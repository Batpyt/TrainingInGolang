package 链表

/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。



示例 1：


输入：head = [1,2,2,1]
输出：true
示例 2：


输入：head = [1,2]
输出：false

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
递归对比两边节点的值，从两端往中间递归
*/
func isPalindrome(head *ListNode) bool {
	//left代表左端节点，从head开始
	left := head
	//递归方法
	var check func(node *ListNode) bool
	check = func(cur *ListNode) bool {
		//由于cur从head开始，所以会一直走到尾节点然后直接返回true
		if cur != nil {
			//有递归结果且是false，直接返回
			if !check(cur.Next) {
				return false
			}
			//尾节点返回true之后，这里就会开始对比两端的节点了
			if cur.Val != left.Val {
				return false
			}
			//由于每次递归回来的cur都是从后往前倒序的，所以left也要右移
			left = left.Next
		}
		return true
	}
	return check(head)
}
