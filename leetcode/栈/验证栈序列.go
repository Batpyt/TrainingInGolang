package 栈

/*
给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 。



示例 1：

输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
示例 2：

输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/validate-stack-sequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
这道题需要利用给定的两个数组 \textit{pushed}pushed 和 \textit{popped}popped 的如下性质：

数组 \textit{pushed}pushed 中的元素互不相同；

数组 \textit{popped}popped 和数组 \textit{pushed}pushed 的长度相同；

数组 \textit{popped}popped 是数组 \textit{pushed}pushed 的一个排列。

根据上述性质，可以得到如下结论：

栈内不可能出现重复元素；

如果 \textit{pushed}pushed 和 \textit{popped}popped 是有效的栈操作序列，则经过所有的入栈和出栈操作之后，每个元素各入栈和出栈一次，栈为空。

因此，可以遍历两个数组，模拟入栈和出栈操作，判断两个数组是否为有效的栈操作序列。

模拟入栈操作可以通过遍历数组 \textit{pushed}pushed 实现。由于只有栈顶的元素可以出栈，因此模拟出栈操作需要判断栈顶元素是否与 \textit{popped}popped 的当前元素相同，如果相同则将栈顶元素出栈。由于元素互不相同，因此当栈顶元素与 \textit{popped}popped 的当前元素相同时必须将栈顶元素出栈，否则出栈顺序一定不等于 \textit{popped}popped。

根据上述分析，验证栈序列的模拟做法如下：

遍历数组 \textit{pushed}pushed，将 \textit{pushed}pushed 的每个元素依次入栈；

每次将 \textit{pushed}pushed 的元素入栈之后，如果栈不为空且栈顶元素与 \textit{popped}popped 的当前元素相同，则将栈顶元素出栈，同时遍历数组 \textit{popped}popped，直到栈为空或栈顶元素与 \textit{popped}popped 的当前元素不同。

遍历数组 \textit{pushed}pushed 结束之后，每个元素都按照数组 \textit{pushed}pushed 的顺序入栈一次。如果栈为空，则每个元素都按照数组 \textit{popped}popped 的顺序出栈，返回 \text{true}true。如果栈不为空，则元素不能按照数组 \textit{popped}popped 的顺序出栈，返回 \text{false}false。

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/validate-stack-sequences/solution/yan-zheng-zhan-xu-lie-by-leetcode-soluti-cql0/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	index := 0

	for _, push := range pushed {
		stack = append(stack, push)
		for len(stack) > 0 && stack[len(stack)-1] == popped[index] {
			stack = stack[:len(stack)-1]
			index++
		}
	}
	return len(stack) == 0
}
