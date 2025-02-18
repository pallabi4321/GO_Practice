package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{33, 4, 56, 12, 34, 67, 7}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println(arr)
}
