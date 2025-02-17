package main

//using maps in go , from a string word , take out vowels and print those vowels in reverse order

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "pallabi"
	vowel := "aeiou"

	rvowels := []rune(vowel)

	var srune []rune
	for i := 0; i <= len(rvowels)-1; i++ {
		if strings.Contains(str, string(rvowels[i])) {
			srune = append(srune, rvowels[i])
		}
	}

	sort.Slice(srune, func(i, j int) bool {
		return srune[i] > srune[j]
	})

	fmt.Println("sorted vowels:", string(srune))
}
