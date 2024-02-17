package main

import "fmt"

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

func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(slice []int) {
		dst(slice)
		for _, f := range src {
			f(slice)
		}
	}
}
