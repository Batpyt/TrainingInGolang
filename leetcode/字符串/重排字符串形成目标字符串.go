package main

import "fmt"

//func main() {
//	fmt.Println(rearrangeCharacters("lrnvlcqukanpdnluowenfxquitzryponxsikhciohyostvmkapkfpglzikitwiraqgchxnpryhwpuwpozacjhmwhjvslprqlnxrk",
//		"woijih"))
//}

func rearrangeCharacters(s string, target string) int {
	m := make(map[string]int)
	for _, t := range target {
		temp := string(t)
		if _, ok := m[temp]; !ok {
			m[temp] = 1
		} else {
			m[temp]++
		}
	}

	resM := make(map[string]int)
	for _, t := range s {
		temp := string(t)
		if _, ok := resM[temp]; !ok {
			resM[temp] = 1
		} else {
			resM[temp]++
		}
	}

	fmt.Println(m)
	fmt.Println(resM)

	res := -1
	for k, v := range m {
		resv, ok := resM[k]
		if !ok {
			return 0
		}
		if res < 0 {
			res = resv / v
		}
		res = min(res, resv/v)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
