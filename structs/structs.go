package main

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Triangle struct {
	base   float64
	height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() float64 {
	var flength float64
	flength = 2 * (r.width + r.height)
	return flength
}

func (c Circle) Perimeter() float64 {
	var flength float64
	flength = 2*c.Radius + math.Pi
	return flength
}

func (t Triangle) Perimeter() float64 {
	var flength float64
	flength = 0.0 // not easy - need sin and cos and to deide the type of triangle..
	return flength
}

func (r Rectangle) Area() float64 {
	var flarea float64
	flarea = r.width * r.height
	return flarea
}

func (c Circle) Area() float64 {
	var flarea float64
	flarea = math.Pow(c.Radius, 2) * math.Pi
	return flarea
}

func (t Triangle) Area() float64 {
	var flarea float64
	flarea = 0.5 * t.base * t.height
	return flarea
}
