package helper

import "strings"

func Reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func Uppcase(s string) string {
	return strings.ToUpper(s)
}

func Modify(s string, opt bool) string {
	if opt {
		return s + "_MODIFY"
	} else {
		return s + "_modify"
	}
}
