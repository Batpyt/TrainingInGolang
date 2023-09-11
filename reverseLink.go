package TrainingInGolang

import "fmt"

type LinkedList struct {
	next *LinkedList
	val  int
}

func main() {
	head := &LinkedList{
		next: &LinkedList{
			next: &LinkedList{
				next: &LinkedList{
					next: nil,
					val:  4,
				},
				val: 3,
			},
			val: 2,
		},
		val: 1,
	}

}

func printLink(head *LinkedList) {
	p := head
	for p != nil {
		fmt.Print(p.val, " -> ")
	}
}

/*
p -> 1 -> 2 -> 3 -> 4
p   cur  next

	4 <- 3 <- 2 <- 1
*/
func reverseLink(head *LinkedList) *LinkedList {
	pre := &LinkedList{}
	cur := head
	for cur != nil {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	return pre
}

/*
m = 2, n = 4

	dummy -> 1 -> 2 -> 3 -> 4 -> 5 -> 6
	        pre  lstart    rend  rnext
*/
func reverseMtoNLink(head *LinkedList, m, n int) *LinkedList {
	dummy := &LinkedList{}

	dummy.next = head
	pre := dummy

	for i := 0; i < m-1; i++ {
		pre = pre.next
	}
	leftStart := pre.next

	rightEnd := pre
	for i := 0; i < n-m-1; i++ {
		rightEnd = rightEnd.next
	}
	rightNext := rightEnd.next

	pre.next = nil
	rightEnd.next = nil

	reverseLink(leftStart)

	pre.next = rightEnd
	leftStart.next = rightNext

	return dummy.next
}
