package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Polygon struct {
	Name   string
	Points []Point
}

// Describe the polygon, by name
func Describe(shape Polygon) {
	fmt.Printf("This polygon is a %s\n", shape.Name)
}

// Print instructions on how to draw this polygon on the plane
func DrawInstructions(shape Polygon) {
	fmt.Printf("To make this %s:\n", shape.Name)
	for lcv := 1; lcv < len(shape.Points); lcv++ {
		fmt.Printf("- draw a line from (%0.0f,%0.0f) to (%0.0f,%0.0f)\n",
			shape.Points[lcv-1].X, shape.Points[lcv-1].Y, shape.Points[lcv].X, shape.Points[lcv].Y)
	}
}

// Get the length of the line segment defined by the coordinates of the terminal points
func lineSegmentLength(alpha Point, beta Point) float64 {
	xDeltaSquared := math.Pow(alpha.X-beta.X, 2)
	yDeltaSquared := math.Pow(alpha.Y-beta.Y, 2)
	length := math.Sqrt(xDeltaSquared + yDeltaSquared)

	return length
}

// Print the perimeter of the polygon on the plane
func Function1(shape Polygon) {
	perimeter := 0.0
	for lcv := 1; lcv < len(shape.Points); lcv++ {
		perimeter = perimeter + lineSegmentLength(shape.Points[lcv-1], shape.Points[lcv])
	}
	fmt.Printf("The perimeter of this %s is %0.1f\n", shape.Name, perimeter)
}

// Print the area the polygon covers on the plane
func Function2(shape Polygon) {
	// https://www.mathopenref.com/coordpolygonarea.html
	subtotal := 0.0
	for lcv := 0; lcv < len(shape.Points); lcv++ {
		next := (lcv + 1) % len(shape.Points)
		subtotal += (shape.Points[lcv].X * shape.Points[next].Y) - (shape.Points[lcv].Y * shape.Points[next].X)
	}
	area := math.Abs(subtotal / 2.0)

	fmt.Printf("The area of this %s is %0.1f\n\n", shape.Name, area)
}

func main() {
	// triangle: (0,0), (4,0), (4,4)
	triangle := Polygon{Name: "Triangle", Points: []Point{{0, 0}, {4, 0}, {4, 4}}}
	Describe(triangle) // This polygon is a Triangle
	DrawInstructions(triangle)	// To make this Triangle:
								// - draw a line from (0,0) to (4,0)
								// - draw a line from (4,0) to (4,4)
	Function1(triangle) // The perimeter of this Triangle is 8.0
	Function2(triangle) // The area of this Triangle is 8.0

	// square: (0,0), (4,0), (4,4), (0,4)
	square := Polygon{Name: "Square", Points: []Point{{0, 0}, {4, 0}, {4, 4}, {0, 4}}}
	Describe(square) // This polygon is a Square
	DrawInstructions(square)	// To make this Square:
								// - draw a line from (0,0) to (4,0)
								// - draw a line from (4,0) to (4,4)
								// - draw a line from (4,4) to (0,4)
	Function1(square) // The perimeter of this Square is 12.0
	Function2(square) // The area of this Square is 16.0
}
