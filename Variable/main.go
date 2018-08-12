package main

import (
	"fmt"
	"math"
)

type Object struct {
	name string
}

type Info interface {
	getInfo()
}

type Geometry interface {
	getArea() float64
}

type Square struct {
	// extends
	Object
	width float64
}

// implements interface Geometry
func (s Square) getArea() float64 {
	return s.width * 4
}

func (s Square) getInfo() {
	fmt.Println(s.getArea())
}

type circle struct {
	Object
	radius float64
}

func (c circle) getArea() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) getInfo() {
	fmt.Println(c.getArea())
}

func showArea(g Geometry) {
	fmt.Println(fmt.Sprintf("my area : %f", g.getArea()))
}

func main() {

	obj := Object{name: "someName"}

	s := Square{obj, 20}

	s.getInfo()

	showArea(s)

	circle{radius: 12}.getInfo()
}
