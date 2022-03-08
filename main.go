package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

// Declare structs
type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// Declare a shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Create receiver functions to calculate the area and perimeter
// for each of the declared structs (implement the shape interface)

// Receiver functions to calculate area
func (s rectangle) Area() float64 {
	return s.width * s.height
}

func (s circle) Area() float64 {
	return math.Pi * s.radius * s.radius
}

// Receiver functions to calculate perimeter
func (s rectangle) Perimeter() float64 {
	return 2 * (s.width + s.height)
}

func (s circle) Perimeter() float64 {
	return 2 * math.Pi * s.radius
}

// Create a function that prints the area and perimeter for any shape
func PrintShapeInfo(s Shape) {
	// Use Golang's built-in reflect package to determine the type
	// of Shape interface passed to the function
	formattedType := strings.Title(reflect.TypeOf(s).String()[5:])

	// Format the output to display the type of the shape along with its
	// area and perimeter, rounded to two decimals
	fmt.Printf("%s => Area: %.2f, Perimeter: %.2f\n", formattedType, s.Area(), s.Perimeter())
}

func main() {
	// Initialize rectangle & circle structs
	r := rectangle{width: 4, height: 12}
	c := circle{radius: 3}

	// Print the info of the rectangle
	PrintShapeInfo(r)
	// Print the info of the circle
	PrintShapeInfo(c)
}
