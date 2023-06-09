package 树

/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

示例 1：
输入：root = [1,null,2,3]
输出：[1,3,2]
*/
/*
		1
	  /   \
	2		3
  /   \    /
4	   5   6
*/

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	//维护栈来存放tree节点
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		//一直往节点的左边遍历直到左边没有子节点
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		//将stack中最后一个节点（也就是nil的前一个节点）记为root
		root = stack[len(stack)-1]
		//并把该节点从stack中删除（pop）
		stack = stack[:len(stack)-1]
		//将当前节点的值存入res，因为这已经是最左边的节点，可以记录
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
