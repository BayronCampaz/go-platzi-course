package mypackage

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	base float64
}

type Rectangle struct {
	base   float64
	height float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.base * s.base
}

func (r Rectangle) Area() float64 {
	return r.base * r.height
}

func (c Circle) Area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (s *Square) SetBase(newBase float64) {
	s.base = newBase
}

func (r *Rectangle) SetBase(newBase float64) {
	r.base = newBase
}

func (r *Rectangle) SetHeight(newHeight float64) {
	r.height = newHeight
}

func (c *Circle) SetRadius(newRadius float64) {
	c.radius = newRadius
}
