package main

func lengthOfLongestSubstring(s string) int {
	exist := make(map[string]bool)

	var right, res int
	for left := 0; left < len(s); left++ {
		if left > 0 {
			delete(exist, string(s[left-1]))
		}

		for right < len(s) {
			if _, ok := exist[string(s[right])]; !ok {
				exist[string(s[right])] = true
				right++
			} else {
				break
			}
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}
