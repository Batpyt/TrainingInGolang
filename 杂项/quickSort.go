package main

import "fmt"

//
//func main() {
//	arr := []int{3,1,5,6,4,3,2,2,3,4,5,5,6,6,88,66,3,22,999,54}
//	quickSort2(arr,0, len(arr) - 1)
//	fmt.Println(arr)
//}

func quickSort2(arr []int, low, high int) {
	if low < high {
		keyPoint := helper(arr, low, high)
		helper(arr, low, keyPoint-1)
		helper(arr, keyPoint+1, high)
	}
}

func helper(arr []int, low, high int) int {
	key := arr[low]

	for low < high {
		for low < high && arr[high] >= key {
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] <= key {
			low++
		}
		arr[high] = arr[low]
	}

	arr[low] = key
	fmt.Println(arr)
	return low
}
