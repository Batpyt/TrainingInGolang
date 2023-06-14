package main

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	//两数相乘，结果的长度最大为两数的长度和
	len1, len2 := len(num1), len(num2)
	//提前初始化一个对应长度的数组，num1[i] * num2[j]的值存在arr[i+j+1]
	arr := make([]int, len1+len2)
	for i := len1 - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := len2 - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			arr[i+j+1] += x * y
		}
	}

	//将大于10的某一位数进位
	for i := len1 + len2 - 1; i > 0; i-- {
		arr[i-1] += arr[i] / 10
		arr[i] %= 10
	}

	//存到字符串中，如果第一个数为0，舍弃
	res := ""
	for i := 0; i < len1+len2; i++ {
		if i == 0 && arr[i] == 0 {
			continue
		}
		res += strconv.Itoa(arr[i])
	}
	return res
}
