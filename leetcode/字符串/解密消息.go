package main

import (
	"strings"
)

//func main() {
//	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
//}

func decodeMessage(key string, message string) string {
	decode := make(map[string]string)
	letter := "abcdefghijklmnopqrstuvwxyz"

	key = strings.Replace(key, " ", "", -1)
	index := 0

	for _, k := range key {
		if _, ok := decode[string(k)]; !ok {
			decode[string(k)] = string(letter[index])
			index++
		}
	}

	res := ""
	for _, m := range message {
		if string(m) == " " {
			res += string(m)
		} else {
			res += decode[string(m)]
		}
	}
	return res
}
