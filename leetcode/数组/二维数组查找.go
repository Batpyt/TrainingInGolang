package 数组

/*
在一个 n * m 的二维数组中，每一行都按照从左到右 非递减 的顺序排序，每一列都按照从上到下 非递减 的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。



示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	//纵坐标
	i := 0
	//横坐标
	j := len(matrix[0]) - 1

	//从左上角开始查找
	for j >= 0 && i < len(matrix) {
		//当当前元素大于目标值，应当往左边查找，横坐标--
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target { //当当前元素小于目标值，往下查找，纵坐标++
			i++
		} else if target == matrix[i][j] {
			return true
		}
	}
	return false
}
