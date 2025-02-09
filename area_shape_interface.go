// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

type shape interface {
	area() int
	shapeType() string
}

type square struct {
	s int
}

type rectangle struct {
	a int
	b int
}

func (sq square) area() int {
	return sq.s * sq.s
}
func (sq square) shapeType() string {
	return "square"
}

func (r rectangle) area() int {
	return r.a * r.b
}

func printArea(s shape) {
	fmt.Println("Area Shape:", s.shapeType())
	fmt.Println("Area", s.area())
}

func (r rectangle) shapeType() string {
	return "rectangle"
}

func main() {
	sq := square{s: 5}
	re := rectangle{a: 4, b: 6}

	printArea(sq)
	printArea(re)
}
