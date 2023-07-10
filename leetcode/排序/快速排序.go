package main

import "fmt"

func main() {
	list := []int{4, 2, 1, 5, 67, 7, 4, 2, 12, 13, 1, 331, 312, 5, 6, 656565656, 343, 43, 34, 4}
	//list := []int{5, 1, 9, 9, 2, 3, 4, 6, 7}
	//list := []int{1,2,3,2,1,5,6,7,777,5,6,77744,999999,4,3,111}
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}

/*
快排就是挑选一个基准值，然后将所有比基准值小的放到其左边，大的放到右边。
然后在基准值两侧的子数组递归进行该操作。
*/
func quickSort(list []int, low, high int) {
	if low < high {
		//splitPoint既是基准值数组下标位置，在list[slitPoint]左边的都小于他，右边的都大于它
		splitPoint := partition(list, low, high)
		//递归排序左右子数组
		quickSort(list, low, splitPoint-1)
		quickSort(list, splitPoint+1, high)
	}
}

/*
5,1,9,9,2,3,4,6,7
[4 1 3 2 5 9 9 6 7]
[2 1 3 4 5 9 9 6 7]
[1 2 3 4 5 9 9 6 7]
[1 2 3 4 5 7 9 6 9]
[1 2 3 4 5 6 7 9 9]
[1 2 3 4 5 6 7 9 9]
*/
func partition(list []int, low, high int) int {
	//直接将数组的第一个选为基准值
	key := list[low]

	/*
		以下的操作本质为找出右边小于key的和左边大于key的进行交换。
	*/
	for low < high {
		//向前移动右指针直到找到比基准值小的，然后将左指针位置的值赋值为该值
		for low < high && list[high] >= key {
			high--
		}
		list[low] = list[high]

		//向后移动左指针直到找到比基准值大的，将右指针位置的值赋值为该值，完成了交换
		for low < high && list[low] <= key {
			low++
		}
		list[high] = list[low]
	}

	//经过交换，左指针的左边已没有比key大的，右边已没有比key小的，故将基准值key赋值给左指针位置
	list[low] = key
	fmt.Println(list)
	//返回新的基准下标
	return low
}
