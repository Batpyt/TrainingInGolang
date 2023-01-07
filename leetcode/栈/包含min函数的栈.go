package main

import "sort"

type MinStack struct {
	stack  []int
	sorted []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	sort.Ints(this.sorted)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	sort.Ints(this.sorted)
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.sorted[0]
}
