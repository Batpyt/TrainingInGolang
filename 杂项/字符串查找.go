package main

/*
helloworld
     world
*/

//func main() {
//	fmt.Println(findStr("helloworld", "world"))
//	fmt.Println(findStr("helloworld", "llo"))
//	fmt.Println(findStr("helloworld", "iii"))
//}

func findStr(str, target string) int {
	if len(target) > len(str) {
		return -1
	}

	left, right := 0, len(target)-1
	for right < len(str) {

		temp := str[left : right+1]
		//fmt.Println(temp)
		if temp == target {
			return left
		}
		left++
		right++
	}
	return -1
}
