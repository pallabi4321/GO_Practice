package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "aabbcccccddffe"

	srune := []rune(str)
	clet := make(map[string]int)

	for i := 0; i <= len(srune)-1; i++ {
		if _, ok := clet[string(srune[i])]; ok {
			clet[string(srune[i])] += 1
		} else {
			clet[string(srune[i])] = 1
		}
	}

	var keys []string
	for v := range clet {
		keys = append(keys, v)
	}

	fmt.Println(keys)

	sort.Slice(keys, func(i, j int) bool {
		return clet[keys[i]] > clet[keys[j]]
	})

	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(k, clet[k])
	}
}
