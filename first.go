package main

import "fmt"
import "math"

func main() {

	fmt.Println(a)
}

func MakeDistOrigin (o_x, o_y float64)
	func (float64, float64) float64 {
		fn:= func (x, y float64) float64 {
			return math.Sqrt (math.Pow(x - o_x, 2) + 
		math.Pow(y - o_y, 2))
		}
		return fn
	}
