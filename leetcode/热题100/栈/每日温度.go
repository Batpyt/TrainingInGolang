package main

/*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。



示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
示例 2:

输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
示例 3:

输入: temperatures = [30,60,90]
输出: [1,1,0]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/daily-temperatures
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
以下用一个具体的例子帮助读者理解单调栈。对于温度列表
[73,74,75,71,69,72,76,73]，单调栈stack 的初始状态为空，答案ans 的初始状态是
[0,0,0,0,0,0,0,0]，按照以下步骤更新单调栈和答案，其中单调栈内的元素都是下标，括号内的数字表示下标在温度列表中对应的温度。

当i=0 时，单调栈为空，因此将0 进栈。
stack=[0(73)]
ans=[0,0,0,0,0,0,0,0]

当i=1 时，由于74 大于73,因此移除栈顶元素0，赋值
ans[0]:=1−0，将1 进栈。
stack=[1(74)]
ans=[1,0,0,0,0,0,0,0]

i=2 时，由于75 大于74，因此移除栈顶元素1，赋值
ans[1]:=2−1，将2 进栈。
stack=[2(75)]
ans=[1,1,0,0,0,0,0,0]

当i=3 时，由于71 小于75，因此将3 进栈。
stack=[2(75),3(71)]
ans=[1,1,0,0,0,0,0,0]

当i=4 时，由于69 小于71，因此将4 进栈。
stack=[2(75),3(71),4(69)]
ans=[1,1,0,0,0,0,0,0]

当i=5 时，由于72 大于69 和71，因此依次移除栈顶元素4 和3，
赋值ans[4]:=5−4 和 ans[3]:=5−3，将5 进栈。
stack=[2(75),5(72)]
ans=[1,1,0,2,1,0,0,0]

当i=6 时，由于76 大于72 和75，因此依次移除栈顶元素5 和2，赋值
ans[5]:=6−5 和
ans[2]:=6−2，将6 进栈。
stack=[6(76)]
ans=[1,1,4,2,1,1,0,0]

当i=7 时，由于73 小于76，因此将7 进栈。
stack=[6(76),7(73)]
ans=[1,1,4,2,1,1,0,0]
*/
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	//栈中存放的是温度数组的下标
	stack := []int{}

	for i := 0; i < len(temperatures); i++ {
		temp := temperatures[i]

		//当当前遍历到的温度 > 栈顶下标对应的温度时，说明遍历到了对于栈顶下标来说的下一个更高的温度
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			//pop
			stack = stack[:len(stack)-1]
			//当前栈顶下标对应的下一个更高温度出现在几天后 = 当前温度下标i - 当前栈顶下标 =
			res[index] = i - index
		}
		//无论如何，当前温度下标都入栈
		stack = append(stack, i)
	}
	return res
}
