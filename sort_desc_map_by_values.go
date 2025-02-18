package main

import (
	"fmt"
	"sort"
)

func findLetterCount(str string) map[string]int {
	rstr := []rune(str)

	countLet := make(map[string]int)
	for _, val := range rstr {
		if _, ok := countLet[string(val)]; ok {
			countLet[string(val)] += 1
		} else {
			countLet[string(val)] = 1
		}
	}

	return countLet
}

func printInDesc(letCount map[string]int) {
	var keyValues []string
	for key := range letCount {
		keyValues = append(keyValues, key)
	}

	sort.Strings(keyValues)
	sort.Slice(keyValues, func(i, j int) bool {
		return letCount[keyValues[i]] > letCount[keyValues[j]]
	})

	for ind := range keyValues {
		fmt.Println(keyValues[ind], letCount[keyValues[ind]])
	}
}

func main() {
	str := "hhhndfffujjsnnn"

	letCount := findLetterCount(str)
	printInDesc(letCount)
}
