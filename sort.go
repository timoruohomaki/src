/* INSTRUCTIONS

Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

*/

package main

import "fmt"

func BubbleSort(sli []int) {

	for i := 0; i < len(sli)-1; i++ {
		Swap(sli, i)
	}

}

func Swap(sli2swap []int, idx int) {
	for j := 0; j < len(sli2swap)-idx-1; j++ {
		if sli2swap[j] > sli2swap[j+1] {
			sli2swap[j], sli2swap[j+1] = sli2swap[j+1], sli2swap[j]
		}
	}

}

func main() {

	fmt.Println("Please enter the integers to be sorted:")

	a := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Println("Enter value ", i+1, "of 10:")
		fmt.Scanln(&a[i])
	}

	BubbleSort(a)

	fmt.Println("Your integers sorted by value:")
	fmt.Println(a)
}
