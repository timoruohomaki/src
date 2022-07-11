/* INSTRUCTIONS

Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))

*/

package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(o_a, o_v, o_s float64) func(float64) float64 {
	return func(t float64) float64 {
		return ((0.5 * o_a * math.Pow(t, 2)) + (o_v * t) + o_s)
	}
}

func main() {

	// ask for the values and put them into slice a

	a := make([]float64, 4)

	fmt.Println("** Displacement calculator **")
	fmt.Println("-----------------------------")

	fmt.Println("Enter value for acceleration [a]: ")
	fmt.Scanln(&a[0])

	fmt.Println("Enter value for initial velocity [v0]: ")
	fmt.Scanln(&a[1])

	fmt.Println("Enter value for initial displacement [s0]: ")
	fmt.Scanln(&a[2])

	fmt.Println("Enter value for time of movement in seconds [t]: ")
	fmt.Scanln(&a[3])

	// call calculator function with provided input

	fn := GenDisplaceFn(a[0], a[1], a[2])

	// print results

	fmt.Println("RESULT: Displacement after", a[3], "seconds is", fn(a[3]), "meters.")

}
