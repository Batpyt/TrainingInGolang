package main

import (
	"fmt"
	"math"
	"strconv"
)

//func main() {
//	reader := bufio.NewReader(os.Stdin)
//
//	str, _, _ := reader.ReadLine()
//	strList := strings.Split(string(str), " ")
//
//	list := make([]int, 0)
//	for _, s := range strList {
//		temp, _ := strconv.ParseInt(s, 10, 10)
//		list = append(list, int(temp))
//	}
//
//	fmt.Println(solveN(list))
//}

func solveN(list []int) int {
	initAve := average(list)
	if initAve == 128 {
		return 0
	}
	//fmt.Println(list)
	//return 0
	newGap := math.Abs(initAve - 128)
	if initAve < 128 {
		return increase(list, 0, newGap)
	}
	return decrease(list, 0, newGap)
}

func increase(list []int, count int, lastGap float64) int {
	newList := make([]int, 0)
	for _, l := range list {
		if l+1 > 255 {
			newList = append(newList, 255)
		} else {
			newList = append(newList, l+1)
		}
	}
	count++
	ave := average(newList)
	if ave == 128 {
		return count
	}

	newGap := math.Abs(ave - 128)
	if newGap == lastGap {
		return count - 1
	}
	if newGap > lastGap {
		return count - 1
	}
	return increase(newList, count, newGap)
}

func decrease(list []int, count int, lastGap float64) int {
	newList := make([]int, 0)
	for _, l := range list {
		if l-1 < 0 {
			newList = append(newList, 0)
		} else {
			newList = append(newList, l-1)
		}
	}
	count--
	ave := average(newList)

	if ave == 128 {
		return count
	}
	newGap := math.Abs(ave - 128)
	//fmt.Println(newList)
	//fmt.Println(ave)
	//fmt.Println(newGap, lastGap)
	if newGap == lastGap {
		return count
	}
	if newGap > lastGap {
		return count + 1
	}
	return decrease(newList, count, newGap)
}

func average(list []int) float64 {
	sum := 0
	for _, l := range list {
		sum += l
	}

	res, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", float64(sum)/float64(len(list))), 64) // 保留5位小数
	return res
}
