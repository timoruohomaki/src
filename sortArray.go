/*

Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func sortSlice(c chan []int, sli []int) {

	defer wg.Done()

	sortSli := make([]int, len(sli))
	copy(sortSli, sli)
	sort.Ints(sortSli)

	c <- sortSli

}

func main() {

	fmt.Println("=== ARRAY SORTER ===")
	fmt.Println("Enter a series of integers separated by space:")
	fmt.Println("(Enter at least 4 integers.)")

	// get user input and create int arrays

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	inputIntStr := strings.Split(s, " ")

	inputInt := make([]int, len(inputIntStr))

	for i := range inputIntStr {
		j, _ := strconv.Atoi(inputIntStr[i])
		inputInt[i] = j
	}

	// arrange into four parts

	partSize := int(math.Floor(float64(len(inputInt)) / 4))

	q1 := inputInt[:(1 * partSize)]
	q2 := inputInt[(1 * partSize):(2 * partSize)]
	q3 := inputInt[(2 * partSize):(3 * partSize)]
	q4 := inputInt[(3 * partSize):]

	// sort each part separately

	c := make(chan []int, 10)

	wg.Add(4)

	go sortSlice(c, q1)

	go sortSlice(c, q2)

	go sortSlice(c, q3)

	go sortSlice(c, q4)

	wg.Wait()

	close(c)

	sortedInt := make([]int, 0)

	fmt.Println("Partial sorts:")

	for cout := range c {

		fmt.Println(cout)

		sout := make([]int, len(cout))
		copy(sout, cout)
		sortedInt = append(sortedInt, sout...)

	}

	// merge all parts and sort it all

	sort.Ints(sortedInt)

	fmt.Println("Compiled sort:")
	fmt.Println(sortedInt)

}
