package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
0 1 2 3 4 5 6 7 8 9 10 11
			0
		1	    	2
	3     4   	   5  6
7    8  9   10    11
*/

func main() {
	trees := tree([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	preorder(trees)
}

func tree(list []int) *TreeNode {
	treeNodes := make([]*TreeNode, 0)
	for _, l := range list {
		treeNodes = append(treeNodes, &TreeNode{
			Val:   l,
			Left:  nil,
			Right: nil,
		})
	}

	index := 0
	//visited := make(map[int]bool)
	for _, t := range treeNodes {
		if index*2+1 >= len(treeNodes) {
			break
		}
		t.Left = treeNodes[index*2+1]
		if index*2+2 >= len(treeNodes) {
			break
		}
		t.Right = treeNodes[index*2+2]
		index++
	}

	return treeNodes[0]
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(fmt.Sprintf("%v ", root.Val))
	preorder(root.Left)
	preorder(root.Right)
}

func levelOrder(root *TreeNode) {

}
