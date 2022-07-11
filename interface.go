package main

import (
	"fmt"
)

// example of using interface

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle {...}

type Rectangle {...}

func FitInYard(s Shape2D) bool {
	if (s.Area() > 100 &&
		s.Perimeter() > 100) {
			return true
		}
		return false
}