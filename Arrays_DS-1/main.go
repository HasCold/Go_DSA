package main

import (
	"fmt"
	"math"
)

// Data Structure :- In programming, the data we have to handle is quite in a very efficient way and data must be stored in a structured, organized and manageable way.
// Data can be present in a linear and non-linear form
// Algorithm means to perform some efficient operations to retreive the data, save and update data.

// -------------------------------- Arrays -----------------------------------
// 1. Arrays have same data-type.
// 2. Contiguous in memory (Data will be stored at a same time)
//							4-bytes,4-bytes,4-bytes,4-bytes  // Integer ko 4-bytes ki memory allot ki jati ha in 32-bit CPU
//	 var marks [5]int  -->> [int, int, int, int, int]
//	Memory Address Points :  100  104  108  112	 116

// 3. Arrays are linear DS.

func smallestNumFind(arr [6]int) (int, int) {
	smallestInt := math.MaxInt // to start infinite numbers
	var smallestIndex int
	// longInt := math.MinInt

	for i := 0; i < len(arr); i++ {
		if arr[i] < smallestInt {
			smallestInt = arr[i]
			smallestIndex = i
		}
	}

	minNum := math.Min(float64(smallestInt), float64(arr[5]))
	fmt.Println("Shortest way to find the min number ", minNum)

	return smallestInt, smallestIndex
}

func changeArr(arr *[6]int) { // dereferencing the array or array points to the int[6]
	for i := 0; i < len(arr); i++ {
		arr[i] = 2 * arr[i] // Now anything we can do in the array will reflect into the original one's because we passed the array by reference not by passing the value.
	}
}

func reverseArr(arr *[6]int) *[6]int { // dereferencing the array or array points to the int[6]

	start := 0          // index start
	end := len(arr) - 1 // end index

	for start <= end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}

	return arr
}

func swap(arr [6]int) [6]int {
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

		if i == (len(arr) - 1) {
			arr[minIndex] = maxInt
			arr[maxIndex] = minInt
		}
	}

	return arr
}

func main() {
	arr := [...]int{5, 15, 22, 1, -15, 24} // [...] it defines the fixed length of an array by the spread operator when we initilize into the index operator []

	// Find the Smallet no. in the given array of integers
	num, index := smallestNumFind(arr)
	fmt.Println("The smallest number from the given array and index is that :- ", num, index)

	fmt.Println("------------------- Swap Max and Min No. of Array ----------------")
	swapMinAndMaxNum := swap(arr)
	fmt.Println("The Min and Max Swaped Numbers :- ", swapMinAndMaxNum)

	fmt.Println("------------------- Reverse an Actual Array by 2 Pointer Approach ----------------")
	reversedArr := reverseArr(&arr)
	fmt.Println("The reversed array we get :- ", *reversedArr)

	fmt.Println("------------------- Array Pass by reference ----------------")
	// Reference  -->> Pass the array actual memory address location rather than passing the array value
	// In Go, arrays are not passed by reference implicitly. Instead, arrays are passed by value by default.
	changeArr(&arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}
