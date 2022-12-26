package main

func main() {

}

/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。





示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
回溯+枝剪
每一个元素都遍历上下左右的元素来看是否与word相符，若不相符，则返回。
*/

/*
创建控制方向的数组，上下左右，用来给元素遍历下一个
*/
type pair struct {
	x, y int
}

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	//字母不允许重复使用，所以记录遍历过的字母
	visited := make([][]bool, len(board))
	for i, _ := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	//检查当前元素是否与当前位置的word相等，并对board进行遍历的方法。
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		//不相等，直接返回
		if board[i][j] != word[k] {
			return false
		}
		//已是最后一个元素，返回true
		if k == len(word)-1 {
			return true
		}
		//记录当前已遍历
		visited[i][j] = true
		//对当前元素的上下左右开始遍历
		for _, dir := range directions {
			newI := i + dir.x
			newJ := j + dir.y
			if 0 <= newI && newI < len(board) && 0 <= newJ && newJ < len(board[0]) && !visited[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		//这一层遍历没有符合要求的元素，返回false给上一层，并将记录是否遍历过的回溯
		visited[i][j] = false
		return false
	}

	for i, b := range board {
		for j, _ := range b {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
