package main

import (
	"fmt"
	"strings"
	"unicode"
)

//func main() {
//	fmt.Println(strongPasswordCheckerII("IloveLe3tcode!"))
//}

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}

	var lower, upper, number, special bool
	for i := 0; i < len(password); i++ {
		temp := password[i]
		if i < len(password)-1 && password[i] == password[i+1] {
			return false
		}

		if unicode.IsUpper(rune(temp)) {
			fmt.Println("upper")
			upper = true
		}
		if unicode.IsLower(rune(temp)) {
			fmt.Println("lower")
			lower = true
		}
		if unicode.IsNumber(rune(temp)) {
			fmt.Println("number")
			number = true
		}
		if strings.ContainsRune("!@#$%^&*()-+", rune(temp)) {
			special = true
		}
	}
	return lower && upper && number && special
}
