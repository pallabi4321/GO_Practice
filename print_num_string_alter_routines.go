package main

import (
	"fmt"
)

func printNumbers(numChan, strChan chan bool, done chan bool, limit int) {
	for i := 1; i <= limit; i++ {
		<-numChan
		fmt.Println(i)
		strChan <- true
	}
	close(done)
}

func printStrings(numChan, strChan chan bool, strList []string) {
	for i := 0; i < len(strList); i++ {
		<-strChan
		fmt.Println(strList[i])
		numChan <- true
	}
}

func main() {
	limit := 5
	strList := []string{"A", "B", "C", "D", "E"}

	numChan := make(chan bool)
	strChan := make(chan bool)
	done := make(chan bool)

	go printNumbers(numChan, strChan, done, limit)
	go printStrings(numChan, strChan, strList)

	numChan <- true

	<-done
}
