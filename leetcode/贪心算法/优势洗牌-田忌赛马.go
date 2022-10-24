package 贪心算法

import "sort"

/*
给定两个大小相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的数目来描述。

返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。



示例 1：

输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
输出：[2,11,7,15]
示例 2：

输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
输出：[24,32,8,12]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/advantage-shuffle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
我们首先分别将数组 \textit{nums}_1nums
1
​
  和 \textit{nums}_2nums
2
​
  进行排序，随后只需要不断考虑这两个数组的首个元素：

如果 \textit{nums}_1nums
1
​
  的首个元素大于 \textit{nums}_2nums
2
​
  的首个元素，那么就将它们在答案中对应起来，同时从数组中移除这两个元素，并增加一点「优势」；

如果 \textit{nums}_1nums
1
​
  的首个元素小于等于 \textit{nums}_2nums
2
​
  的首个元素，那么移除 \textit{nums}_1nums
1
​
  的首个元素。

当 \textit{nums}_1nums
1
​
  中没有元素时，遍历结束。

这样做的正确性在于：

对于第一种情况，由于 \textit{nums}_1nums
1
​
  是有序的，那么 \textit{nums}_1nums
1
​
  的任意元素大于 \textit{nums}_2nums
2
​
  的首个元素：

如果我们不与 \textit{nums}_2nums
2
​
  的首个元素配对，由于 \textit{nums}_2nums
2
​
  是有序的，之后的元素会更大，这样并不划算；

如果我们与 \textit{nums}_2nums
2
​
  的首个元素配对，我们使用 \textit{nums}_1nums
1
​
  的首个元素，可以使得剩余的元素尽可能大，之后可以获得更多「优势」。

对于第二种情况，由于 \textit{nums}_2nums
2
​
  是有序的，那么 \textit{nums}_1nums
1
​
  的首个元素小于等于 \textit{nums}_2nums
2
​
  中的任意元素，因此 \textit{nums}_1nums
1
​
  的首个元素无法增加任何「优势」，可以直接移除。

在本题中，由于 \textit{nums}_1nums
1
​
  中的每一个元素都要与 \textit{nums}_2nums
2
​
  中的元素配对，而我们是按照顺序考虑 \textit{nums}_2nums
2
​
  中的元素的。因此在遍历结束后，\textit{nums}_2nums
2
​
  中剩余的元素实际上是原先 \textit{nums}_2nums
2
​
  的一个后缀。因此当 \textit{nums}_1nums
1
​
  的首个元素无法配对时，我们给它配对一个 \textit{nums}_2nums
2
​
  的尾元素即可，并将该尾元素移除。

在实际的代码编写中，我们无需真正地「移除」元素。对于 \textit{nums}_1nums
1
​
 ，我们使用一个循环依次遍历其中的每个元素；对于 \textit{nums}_2nums
2
​
 ，我们可以使用双指针 \textit{left}left 和 \textit{right}right。如果 \textit{nums}_1nums
1
​
  的首个元素可以增加「优势」，就配对 \textit{left}left 对应的元素并向右移动一个位置；如果无法配对，就配对 \textit{right}right 对应的元素并向左移动一个位置。

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/advantage-shuffle/solution/you-shi-xi-pai-by-leetcode-solution-sqsf/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	idx1 := make([]int, n)
	idx2 := make([]int, n)
	for i := 1; i < n; i++ {
		idx1[i] = i
		idx2[i] = i
	}

	sort.Slice(idx1, func(i, j int) bool { return nums1[idx1[i]] < nums1[idx1[j]] })
	sort.Slice(idx2, func(i, j int) bool { return nums2[idx2[i]] < nums2[idx2[j]] })

	res := make([]int, n)
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if nums1[idx1[i]] > nums2[idx2[left]] {
			res[idx2[left]] = nums1[idx1[i]]
			left++
		} else {
			res[idx2[right]] = nums1[idx1[i]]
			right--
		}
	}
	return res
}
