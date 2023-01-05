package main

import "fmt"

func main() {
	//fmt.Println(isValid("[]{}()"))
	//fmt.Println(isValid("[]{}{[]()}"))
	fmt.Println(isValid("(])"))
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

/*
if string(c) == "{" || string(c) == "[" || string(c) == "(" {
			stack = append(stack, string(c))
		}
		if string(c) == "}"
*/
