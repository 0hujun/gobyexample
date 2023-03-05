package gobyexample

import (
	"fmt"
	"math"
	"testing"
)

type geometry interface {
	area() float64
	perim() float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return (r.width + r.height) * 2
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return c.radius * math.Pi * 2
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func TestInterface(t *testing.T) {
	r := rect{
		width:  3,
		height: 4,
	}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
