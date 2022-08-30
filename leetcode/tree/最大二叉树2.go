package tree

/*
最大树 定义：一棵树，并满足：其中每个节点的值都大于其子树中的任何其他值。

给你最大树的根节点 root 和一个整数 val 。

就像 之前的问题 那样，给定的树是利用 Construct(a) 例程从列表 a（root = Construct(a)）递归地构建的：

如果 a 为空，返回 null 。
否则，令 a[i] 作为 a 的最大元素。创建一个值为 a[i] 的根节点 root 。
root 的左子树将被构建为 Construct([a[0], a[1], ..., a[i - 1]]) 。
root 的右子树将被构建为 Construct([a[i + 1], a[i + 2], ..., a[a.length - 1]]) 。
返回 root 。
请注意，题目没有直接给出 a ，只是给出一个根节点 root = Construct(a) 。

假设 b 是 a 的副本，并在末尾附加值 val。题目数据保证 b 中的值互不相同。

返回 Construct(b) 。



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-binary-tree-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
如果根节点的值小于给定的整数 \textit{val}val，那么新的树会以 \textit{val}val 作为根节点，并将原来的树作为新的根节点的左子树。

否则，我们从根节点开始不断地向右子节点进行遍历。这是因为，当遍历到的节点的值大于 \textit{val}val 时，由于 \textit{val}val 是新添加的位于数组末尾的元素，那么在构造的结果中，\textit{val}val 一定出现在该节点的右子树中。

当我们遍历到节点 \textit{cur}cur 以及它的父节点 \textit{parent}parent，并且 \textit{cur}cur 节点的值小于 \textit{val}val 时，我们就可以停止遍历，构造一个新的节点，以 \textit{val}val 为值且以 \textit{cur}cur 为左子树。我们将该节点作为 \textit{parent}parent 的新的右节点，并返回根节点作为答案即可。

如果遍历完成之后，仍然没有找到比 \textit{val}val 值小的节点，那么我们构造一个新的节点，以 \textit{val}val 为值，将该节点作为 \textit{parent}parent 的右节点，并返回根节点作为答案即可。

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/maximum-binary-tree-ii/solution/zui-da-er-cha-shu-ii-by-leetcode-solutio-piv2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	for cur := root; cur != nil; cur = cur.Right {
		if val > cur.Val {
			if parent == nil {
				return &TreeNode{
					Val:   val,
					Left:  root,
					Right: nil,
				}
			}
			parent.Right = &TreeNode{
				Val:   val,
				Left:  cur,
				Right: nil,
			}
			return root
		}
		parent = cur
	}
	parent.Right = &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	return root
}
