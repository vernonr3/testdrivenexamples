package main

type Rectangle struct {
	width  float64
	height float64
}

func Perimeter(width, height float64) float64 {
	var flength float64
	flength = 2 * (width + height)
	return flength
}

func Area(width, height float64) float64 {
	var flarea float64
	flarea = width * height
	return flarea
}
