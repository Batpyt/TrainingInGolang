package main

//func main() {
//	list := []int{5,4,3,6,6,3,1,2}
//	quickSort(list, 0, len(list)-1)
//	fmt.Println(list)
//}

func quickSort(list []int, low, high int) {
	if low < high {
		key := sort(list, low, high)
		quickSort(list, low, key-1)
		quickSort(list, key+1, high)
	}
}

/*
 4 5 3 2 6
*/

func sort(list []int, low, high int) int {
	key := list[low]

	for low < high {
		for low < high && list[high] >= key {
			high--
		}
		list[low] = list[high]

		for low < high && list[low] <= key {
			low++
		}
		list[high] = list[low]
	}

	list[low] = key
	return low
}
