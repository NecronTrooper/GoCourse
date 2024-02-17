package main

func main() {
	slice := []int{-5, -10, -1, 5, 15, 1}

	SortSlice(slice)

	IncrementOdd(slice)

	RevereSlice(slice)

	PrintSlice(slice)

	incrementTwiceOddSortPrint := appendFunc(IncrementOdd, IncrementOdd, SortSlice, PrintSlice)

	incrementTwiceOddSortPrint(slice)
}
