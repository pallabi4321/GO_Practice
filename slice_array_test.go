// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	slArr := arr[1:4]
	fmt.Println(arr)
	fmt.Println(slArr)
	fmt.Println(len(arr))
	fmt.Println(len(slArr))
	fmt.Println(cap(slArr))
	fmt.Println("------------")

	slArr = append(slArr, 6)
	fmt.Println(arr)
	fmt.Println(slArr)
	fmt.Println(len(arr))
	fmt.Println(len(slArr))
	fmt.Println(cap(slArr))

	arr = append(arr, 7)
	fmt.Println(arr)
	fmt.Println(slArr)
	fmt.Println(len(arr))
	fmt.Println(len(slArr))
	fmt.Println(cap(slArr))
	fmt.Println("------------")

	slArr = append(slArr, 8)
	fmt.Println(arr)
	fmt.Println(slArr)
	fmt.Println(len(arr))
	fmt.Println(len(slArr))
	fmt.Println(cap(slArr))
	fmt.Println("------------")

}
