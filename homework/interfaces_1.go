package interfaces_1

import (
	"fmt"
)

type figure interface {
	perimeter()
	area()
}
type circle struct {
	radius int
}
type square struct {
	height int
	width  int
}

func (t square) area() {
	fmt.Println(t.height * t.width)
}
func (t square) perimeter() {
	fmt.Println(2*t.height + 2*t.width)
}
func (t circle) area() {
	fmt.Println(3.14 * float64(t.radius))
}
func (t circle) perimeter() {
	fmt.Println(2 * 3.14 * float64(t.radius))
}

func figures(F figure) {
	F.area()
	F.perimeter()
}
func interfaces_1() {

	cirl := circle{2}
	figures(cirl)
	sqr := square{2, 4}
	figures(sqr)
}
