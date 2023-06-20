package main

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”



示例 1：


输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
一直往下寻找，可以理解为后序遍历，当找到符合条件的p、q时，将节点返回，若没找到的话就返回nil
最终左右子树都有返回值的节点就是祖先
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//一直找到了叶子结点还没有p、q，说明这条路不通，返回nil
	if root == nil {
		return nil
	}

	//找到了p/q，返回当前节点，表示这条是通路
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	//递归查找左右子树是否存在通路
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果left 和 right都不为空，说明此时root就是最近公共节点
	// 如果left为空，right不为空，就返回right，说明目标节点是通过right返回的，反之亦然。
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
