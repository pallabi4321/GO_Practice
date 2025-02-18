package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) pop() {
	popped := s.data[len(s.data)-1]
	fmt.Println("Popped:", popped)
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) push(item int) {
	s.data = append(s.data, item)
}

func (s *Stack) print() {
	fmt.Println(s.data)
}

func main() {
	stackArr := Stack{
		data: make([]int, 0),
	}
	fmt.Println("Is stack empty", stackArr.isEmpty())
	fmt.Println("Add item 5:")
	stackArr.push(5)
	stackArr.print()
	fmt.Println("Add item 10:")
	stackArr.push(10)
	stackArr.print()
	fmt.Println("Add item 23:")
	stackArr.push(23)
	stackArr.print()
	stackArr.pop()
	stackArr.print()
	stackArr.pop()
	stackArr.print()
}
