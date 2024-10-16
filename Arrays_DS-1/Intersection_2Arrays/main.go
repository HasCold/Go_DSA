package main

import "fmt"

func intersectionTwoArray(arr1 [5]int, arr2 [4]int) []int {
	uniqueElement := []int{}

	for i := 0; i < len(arr1); i++ {
		isUnique := false
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				isUnique = true
				break
			}
		}

		if isUnique {
			uniqueElement = append(uniqueElement, arr1[i])
		}
	}

	return uniqueElement
}

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [4]int{6, 7, 3, 1}

	// Find common elements between two arrays
	uniqueNums := intersectionTwoArray(arr1, arr2)
	fmt.Println("The elements found in the intersection of two arrays :- ", uniqueNums)
}
