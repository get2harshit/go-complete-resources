package main

import (
	"fmt"
	"math"
)

// Define the interface
type geometry interface {
	area() float64
	perim() float64
}

// Define a struct type: rect
type rect struct {
	width, height float64
}

// Define another struct type: circle
type circle struct {
	radius float64
}

// Implement the geometry interface for rect
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Implement the geometry interface for circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// A generic function that takes any type satisfying the geometry interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Both rect and circle can be passed to measure()
	// because they implicitly implement the geometry interface
	measure(r)
	measure(c)
}
