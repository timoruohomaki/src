package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// access restricted because lower-case letter

func (p Point) DistToOrig() {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {
	p1 := Point(3, 4)
	fmt.Println(p1.DistToOrig())
}
