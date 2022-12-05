package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct { //struct 정의
	width, height float64
}

func (r Rect) area() float64 { //Rect의 area() 메소드
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}
	showArea(r, c)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		fmt.Println(a)
	}
}
