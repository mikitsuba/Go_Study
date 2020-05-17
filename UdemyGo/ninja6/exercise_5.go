package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
}

type CIRCLE struct {
	radius float64
}

type SHAPE interface {
	AREA() float64
}

func (s SQUARE) AREA() float64 {
	area := s.length * s.length
	return area
}

func (c CIRCLE) AREA() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

func INFO(s SHAPE) {
	area := s.AREA()
	fmt.Println(area)
}

func main() {
	sq1 := SQUARE{
		length: 5,
	}

	cir1 := CIRCLE{
		radius: 5,
	}

	INFO(sq1)
	INFO(cir1)
}
