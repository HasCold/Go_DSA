package main

import (
	"fmt"
	"math"
)

func smallestNumFind(arr [8]int) {
	smallestInt := math.MaxInt
	var smallestIndex int

	for i := 0; i < len(arr); i++ {
		if arr[i] < smallestInt {
			smallestInt = arr[i]
			smallestIndex = i
		}
	}

	fmt.Printf("The smallest integer is %v and index is %v from an array \n", smallestInt, smallestIndex)
}

func swapMinAndMax(arr [8]int) {

	minInt := math.MaxInt
	maxInt := math.MinInt

	maxIndex := len(arr) - 1
	minIndex := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < minInt {
			minInt = arr[i]
			minIndex = i
		}

		if arr[i] > maxInt {
			maxInt = arr[i]
			maxIndex = i
		}

		if i == len(arr)-1 {
			arr[minIndex], arr[maxIndex] = maxInt, minInt
		}
	}

	fmt.Println("The Swapped Min and Max Numbers in Array :- ", arr)
}

func reversedArr(arr *[8]int) *[8]int {
	st := 0
	end := len(arr) - 1

	for st < end {
		temp := arr[st]
		arr[st] = arr[end]
		arr[end] = temp

		st++
		end--
	}

	return arr
}

func main() {
	// 5, 15, 22, 1, -15, 24
	arr := [...]int{1, 22, -5, 6, 8, 40, -1, 0} // defines the fixed length of an array by the spread operator into the index operator [].

	// Find the smallest number in the given array.
	smallestNumFind(arr)

	// Swap Min and Max No. of arrays
	swapMinAndMax(arr)

	// Reversed an actual array by two pointer approach
	revArr := reversedArr(&arr)
	fmt.Println("Reversed Array :-", *revArr)

}
