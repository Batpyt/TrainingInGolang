package tree

/*
输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。

假设输入的前序遍历和中序遍历的结果中都不含重复的数字。



示例 1:


Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
示例 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
前序遍历：根左右
中序遍历：左根右
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	//前序遍历的第一个元素即为根结点
	root := &TreeNode{preorder[0], nil, nil}

	//定位到根节点在中序遍历数组中的位置
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	// 递归地构造左子树，并连接到根节点
	// preorder[1: len(inorder[:i])+1] 为前序遍历的【左】，inorder[:i] 为中序遍历的【左】
	//其中，【len(inorder[:i])+1】是根据中序遍历根的位置来计算左子树的长度
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])

	// 递归地构造右子树，并连接到根节点
	// 这两个数组为对应的两个【右】
	//其中，len(inorder[:i])+1: 是根据中序遍历根的位置来计算右子树的长度
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
