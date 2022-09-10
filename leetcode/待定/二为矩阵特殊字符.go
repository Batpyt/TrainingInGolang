package main

/*
给你一个大小为 rows x cols 的矩阵 mat，其中 mat[i][j] 是 0 或 1，请返回 矩阵 mat 中特殊位置的数目 。

特殊位置 定义：如果 mat[i][j] == 1 并且第 i 行和第 j 列中的所有其他元素均为 0（行和列的下标均 从 0 开始 ），则位置 (i, j) 被称为特殊位置。



示例 1：

输入：mat = [[1,0,0],
            [0,0,1],
            [1,0,0]]
输出：1
解释：(1,2) 是一个特殊位置，因为 mat[1][2] == 1 且所处的行和列上所有其他元素都是 0
示例 2：

输入：mat = [[1,0,0],
            [0,1,0],
            [0,0,1]]
输出：3
解释：(0,0), (1,1) 和 (2,2) 都是特殊位置

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/special-positions-in-a-binary-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numSpecial(mat [][]int) int {
	res := 0
	count := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			count = 0
			if mat[i][j] != 1 {
				continue
			}
			for zong := 0; zong < len(mat); zong++ {
				if mat[zong][j] == 1 && zong != i {
					count++
				}
			}
			for heng := 0; heng < len(mat[i]); heng++ {
				if mat[i][heng] == 1 && heng != j {
					count++
				}
			}
			if count == 0 {
				res++
			}
		}
	}
	return res
}

//func main() {
//	fmt.Println(numSpecial([][]int{{1,0,0},{0,1,0},{0,0,1}}))
//}
