package main

func replaceSpace(s string) string {
	var res string
	for _, ss := range s {
		if string(ss) == " " {
			res += "%20"
		} else {
			res += string(ss)
		}
	}
	return res
}
