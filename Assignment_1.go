package main

import "fmt"

/*
Assignment #1:
1. Create a private repository in github.com for the Assignment. Add tmustakhov as a Collaborator to your repo.

2. Write a program and implement the following functions:

	func SortSlice(slice []int) - a function to sort a slice. It's prohibited to use sort package.

	func IncrementOdd(slice []int) - a function that increments integers by one in odd positions(1,3,5,...)

	func PrintSlice(slice []int) - a function that prints the slice

	func RevereSlice(slice []int) - a function that reverses the slice

3.  Implement a function of the following type:
		appendFunc(dst func([]int), src ...func([[int)) func([]int)
 	This function takes as an argument some function for processing slices dst and an unlimited number of other
  	processing functions that need to be “attached(added)” to the dst function and return a new function
*/

func SortSlice(slice []int) { //bubble sort
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func IncrementOdd(slice []int) { //slice []int{0,1(+1),2,3(+1),...}
	n := len(slice)
	for i := 0; i < n-1; i++ {
		if i%2 == 1 {
			slice[i] = slice[i] + 1
		}
	}
}

func PrintSlice(slice []int) { // prints with "," after every element
	n := len(slice)
	fmt.Print("[")
	for i := 0; i < n; i++ {
		if n-1 != i {
			fmt.Print(slice[i], ",")
		} else {
			fmt.Print(slice[i])
		}
	}
	fmt.Println("]")
}

func RevereSlice(slice []int) {
	n := len(slice) / 2
	m := len(slice)
	for i := 0; i < n; i++ {
		temp := slice[i]
		slice[i] = slice[m-1-i]
		slice[m-1-i] = temp
	}
}
