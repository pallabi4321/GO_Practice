package main

import (
	"fmt"
)

func reverseString(str string) string {
	strc := []rune(str)
	var rstrc []rune
	for i := len(strc) - 1; i >= 0; i-- {
		rstrc = append(rstrc, strc[i])
	}

	if str == string(rstrc) {
		return "Palindrome"
	}

	return "Not Palindrome"
}

func main() {
	str := "mom"

	fmt.Println(reverseString(str))
}
