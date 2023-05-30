package main

//func main() {}

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。



示例 1：


输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：


输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
维护一个visited二维数组
维护一个方向数组，存放右，下，左，上的元素运算规则，还有一个标记当前是哪个方向的元素
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rows, cols := len(matrix), len(matrix[0])

	//visited二维数组
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	//当记录的原素==total个时，可以结束循环
	total := rows * cols
	//从第一个元素开始
	row, col := 0, 0
	//按右，下，左，上的顺序，为元素下标的移动方向。
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	//从向右移动开始
	dirsIndex := 0

	res := make([]int, 0)

	for i := 0; i < total; i++ {
		//记录当前元素并记为visited
		res = append(res, matrix[row][col])
		visited[row][col] = true

		//判断当前方向的下一个元素是否合法：横纵坐标不能小于0或超过范围，且不能已是visited过的
		nextRow, nextCol := row+dirs[dirsIndex][0], col+dirs[dirsIndex][1]
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols || visited[nextRow][nextCol] {
			//如果下一个不合法，则按顺序换下一个方向，这个计算公式可让其在0，1，2，3内循环
			dirsIndex = (dirsIndex + 1) % 4
		}
		//记录下一个元素的下标
		row += dirs[dirsIndex][0]
		col += dirs[dirsIndex][1]
	}
	return res
}
