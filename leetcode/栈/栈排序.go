package main

/*
栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。该栈支持如下操作：push、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。

示例1:

 输入：
["SortedStack", "push", "push", "peek", "pop", "peek"]
[[], [1], [2], [], [], []]
 输出：
[null,null,null,1,null,2]
示例2:

 输入：
["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
[[], [], [], [1], [], []]
 输出：
[null,null,null,null,null,true]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sort-of-stacks-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type SortedStack struct {
	sortedList []int
}

func SortedStackConstructor() SortedStack {
	return SortedStack{sortedList: []int{}}
}

/*
维护一个单调递增队列即可
*/
func (this *SortedStack) Push(val int) {
	//若队列为空或者插入值为最大，直接插入
	if this.IsEmpty() || val >= this.sortedList[len(this.sortedList)-1] {
		this.sortedList = append(this.sortedList, val)
		return
	}

	//定位到该值应该插入的位置
	i := len(this.sortedList)
	for i > 0 && this.sortedList[i-1] > val {
		i--
	}

	this.sortedList = append(this.sortedList[:i], append([]int{val}, this.sortedList[i:]...)...)
	return
}

func (this *SortedStack) Pop() {
	if !this.IsEmpty() {
		this.sortedList = this.sortedList[1:]
	}
}

func (this *SortedStack) Peek() int {
	if !this.IsEmpty() {
		return this.sortedList[0]
	}
	return -1
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.sortedList) == 0
}
