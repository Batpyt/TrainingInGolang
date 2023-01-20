package main

//func main() {
//	fmt.Println(prefixCount([]string{"pay", "attention", "practice", "attend"}, "at"))
//}

func prefixCount(words []string, pref string) int {
	res := 0
	for _, s := range words {
		if len(s) < len(pref) {
			continue
		}
		if s[:len(pref)] == pref {
			res++
		}
	}
	return res
}
