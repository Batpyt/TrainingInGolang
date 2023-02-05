package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str, _, _ := reader.ReadLine()
	keys := strings.Split(string(str), " ")
	if len(keys) <= 1 {
		fmt.Println("")
		return
	}

	res := solve(keys)
	fmt.Println(res)
}

func solve(keys []string) string {
	keyMap := make(map[string]bool)

	for _, k := range keys {
		keyMap[k] = true
	}

	var res string
	for _, k := range keys {
		if len(k) <= 1 {
			continue
		}
		temp := k
		for true {
			temp = temp[:len(temp)-1]
			if _, ok := keyMap[temp]; !ok {
				break
			}
			if len(temp) == 1 {
				if res == "" {
					res = k
				} else {
					if k > res {
						res = k
					}
				}
			}
		}
	}
	return res
}
