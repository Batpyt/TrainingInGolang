package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
}

func isValid(s string) bool {
	stack := make([]string, 0)
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, string(c))
			continue
		}
		//fmt.Println(stack)

		if len(stack) > 0 {
			if c == ')' && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
				continue
			}
			if c == ']' && stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
				continue
			}
			if c == '}' && stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
				continue
			}
		}

		return false
	}
	return len(stack) == 0
}
