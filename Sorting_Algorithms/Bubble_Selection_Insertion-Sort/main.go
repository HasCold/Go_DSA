//  Here we can see the three sorting algorithms which is of O(n^2)(Quadratic) Time Complexity

// ### 1. **Basic Sorting Algorithms**
// These are easy to understand but not always efficient for large datasets.
//    - **Bubble Sort**: Swaps adjacent elements if they’re in the wrong order. Good for small datasets but has poor performance (O(n²) time complexity).
//    - **Selection Sort**: Repeatedly selects the smallest element from the unsorted part and places it at the beginning. Also O(n²) time complexity.
//    - **Insertion Sort**: Builds the sorted array one item at a time. Suitable for small or nearly sorted datasets, with O(n²) time complexity in the worst case.

package main

import "fmt"

func bubbleSort(arr []int) []int { // Time Complexity O(n^2)
	isSwap := false

	for i := 0; i < len(arr); i++ {
		for j := 0; j < (len(arr)-i)-1; j++ {
			if arr[j] > arr[j+1] {
				isSwap = true
				arr[j], arr[j+1] = arr[j+1], arr[j] // So, arr[j] gets the value of arr[j+1], and arr[j+1] gets the value of arr[j], effectively swapping them.
			}
		}

		// Optimized Bubble Sort
		if !isSwap {
			break
		}
	}

	return arr
}

// Selection Sort Intution :-
// [sorted, unsorted]  // so we will take out the smallest number from the unsorted array and put into sorted order
// So by-default we assume the first element was our smallest element.

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		smallestIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallestIndex] {
				smallestIndex = j
			}
		}
		arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]
		fmt.Println(arr)
	}

	return arr
}

func insertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		prev := i - 1

		for prev >= 0 && arr[prev] > curr {
			arr[prev+1] = arr[prev]
			prev--
		}

		arr[prev+1] = curr // Placing the current element in the correct position
	}

	return arr
}

func main() {
	arr := []int{4, 1, 5, 2, 3}
	// arr := []int{1, 2, 4, 8, 9}
	fmt.Println("The Ascending order implementation of bubble sort algo :- ", bubbleSort(arr))

	arr2 := []int{4, 1, 5, 2, 3}
	fmt.Println("The Ascending order implementation of selection sort algo :- ", selectionSort(arr2))

	arr3 := []int{4, 1, 5, 2, 3}
	fmt.Println("The Ascending order implementation of insertion sort algo :- ", insertionSort(arr3))
}
