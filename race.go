/*
RACE CONDITION EXAMPLE (WEEK 2)

Calling function adder 100 times from two different goroutines.
With goroutines, counter is never 200 as it would be without race condition
but will get arbitrary numbers between about 180 and 195.

Test by removing the go keywords and you will get 200.

Also test by running go run -race race.go to verify the race condition.

*/

package main

import "fmt"

var count int
var calls int

func adder() {
	count++
}

func main() {

	for i := 0; i < 100; i++ {

		go adder()
		go adder()

		// adder()
		// adder()

		calls++

	}

	if count != calls {
		fmt.Println("Count is", count, "after", calls, "calls.")
	}

}
