/*
RACE CONDITION EXAMPLE (WEEK 2)

Calling function adder 100 times from two different goroutines.
With goroutines, counter is never 200 as it would be without race condition
but will get arbitrary numbers between about 180 and 195.

Test by removing the go keywords and you will get 200.
You can also fix the issue by enabling the WaitGroup.

Also test by running go run -race race.go to verify the race condition.

*/

package main

import (
	"fmt"
	"sync"
)

var calls, count int

var wg sync.WaitGroup

func adder() {
	count++
	wg.Done()
}

func main() {

	for i := 0; i < 100; i++ {

		wg.Add(2)
		go adder()
		go adder()

		wg.Wait()

		calls++

	}

	if count != calls {
		fmt.Println("Count is", count, "after", calls, "calls.")
	}

}
