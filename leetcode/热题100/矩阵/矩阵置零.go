package main

/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。



示例 1：


输入：matrix =
[[1,1,1],
 [1,0,1],
 [1,1,1]]
输出：
[[1,0,1],
 [0,0,0],
 [1,0,1]]
示例 2：


输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/set-matrix-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
我们可以用矩阵的第一行和第一列代替方法一中的两个标记数组，以达到O(1) 的额外空间。
但这样会导致原数组的第一行和第一列被修改，无法记录它们是否原本包含0
因此我们需要额外使用两个标记变量分别记录第一行和第一列是否原本包含0
在实际代码中，我们首先预处理出两个标记变量，接着使用其他行与列去处理第一行与第一列，
然后反过来使用第一行与第一列去更新其他行与列，最后使用两个标记变量更新第一行与第一列即可。

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/set-matrix-zeroes/solution/ju-zhen-zhi-ling-by-leetcode-solution-9ll7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

//func main() {
//	matrix := [][]int{{1, 1, 0}, {1, 0, 1}, {1, 1, 1}}
//	setZeroes(matrix)
//	fmt.Println(matrix)
//}

/*
我们可以用矩阵的第一行和第一列代替方法一中的两个标记数组，以达到O(1) 的额外空间。
但这样会导致原数组的第一行和第一列被修改，无法记录它们是否原本包含0
因此我们需要额外使用两个标记变量分别记录第一行和第一列是否原本包含0
在实际代码中，我们首先预处理出两个标记变量，接着使用其他行与列去处理第一行与第一列，
然后反过来使用第一行与第一列去更新其他行与列，最后使用两个标记变量更新第一行与第一列即可。
*/
func setZeroes(matrix [][]int) {
	//row0记录第一行有没有0，col0记录第一列第一个值有没有0
	/*
		[[1,1,0],
		 [1,0,1],
		 [1,1,1]]

		row0 = true
		col0 = false
	*/
	row0, col0 := false, false
	rowLen := len(matrix[0])
	colLen := len(matrix)

	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}

	for _, v := range matrix {
		if v[0] == 0 {
			col0 = true
			break
		}
	}

	//然后从第二行第二列开始遍历，如果遍历到0值，就把其对应的第一行第一列的元素变为0
	/*
		[[1,0,0],
		 [0,0,1],
		 [1,1,1]]
	*/
	for i := 1; i < colLen; i++ {
		for j := 1; j < rowLen; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	/*
		然后将第一行/第一列为0的同行/列的元素变为0
		[[1,0,0],
		 [0,0,0],
		 [1,0,0]]
	*/
	for i := 1; i < colLen; i++ {
		for j := 1; j < rowLen; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	/*
		最终，如果根据一开始的row0，col0记录，将第一行/第一列按情况变为0
		[[0,0,0],
		 [0,0,0],
		 [1,0,0]]
	*/
	if row0 {
		for j := 0; j < rowLen; j++ {
			matrix[0][j] = 0
		}
	}

	if col0 {
		for i := 0; i < colLen; i++ {
			matrix[i][0] = 0
		}
	}
}
