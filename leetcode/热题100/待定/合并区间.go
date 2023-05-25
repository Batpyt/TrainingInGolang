package 待定

import "sort"

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。



示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
先针对第一个元素升序排列，然后对比 inter[i][1]和inter[i+10][0],如果前者小于后者，说明没有重合，
否则重合了，并且记录inter[i][1], inter[i+1][1]的更大值
*/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := make([][]int, 0)
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if pre[1] < intervals[i][0] {
			res = append(res, pre)
			pre = intervals[i]
		} else {
			pre[1] = max(pre[1], intervals[i][1])
		}
	}
	res = append(res, pre)
	return res
}
