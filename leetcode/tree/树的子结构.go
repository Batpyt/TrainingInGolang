package tree

/*
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

输入：A = [1,2,3], B = [3,1]
输出：false
示例 2：

输入：A = [3,4,5,1,2], B = [4,1]
输出：true

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	//分别判断A，A的左，A的右和B是不是相等的，并递归往下判断
	return isSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

/*
辅助方法，负责从当前的a，b节点遍历判断是否a，b是相同结构。
*/
func isSub(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isSub(a.Left, b.Left) && isSub(a.Right, b.Right)
}
