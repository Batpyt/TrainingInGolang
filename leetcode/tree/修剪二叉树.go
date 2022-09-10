package tree

/*
给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。

所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/trim-a-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
对根结点 \textit{root}root 进行深度优先遍历。对于当前访问的结点，如果结点为空结点，直接返回空结点；如果结点的值小于 \textit{low}low，那么说明该结点及它的左子树都不符合要求，我们返回对它的右结点进行修剪后的结果；如果结点的值大于 \textit{high}high，那么说明该结点及它的右子树都不符合要求，我们返回对它的左子树进行修剪后的结果；如果结点的值位于区间 [\textit{low}, \textit{high}][low,high]，我们将结点的左结点设为对它的左子树修剪后的结果，右结点设为对它的右子树进行修剪后的结果。

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/trim-a-binary-search-tree/solution/xiu-jian-er-cha-sou-suo-shu-by-leetcode-qe7q1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
