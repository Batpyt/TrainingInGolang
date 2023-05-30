package main

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/search-a-2d-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/search-a-2d-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
从右上角开始遍历
若 target < i，则target在该列之前
若 target > i，则target在该行之后
超出边界则说明不存在
*/
func searchMatrix(matrix [][]int, target int) bool {
	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		if target == matrix[row][col] {
			return true
		}
		if target < matrix[row][col] {
			col--
		} else {
			row++
		}
	}
	return false
}
