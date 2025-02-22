package main

import "fmt"

func printUniqueValues(arr [8]int) []int {
	var uniqueValues []int

	for i := 0; i < len(arr); i++ {
		isUnique := true

		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == arr[j] {
				isUnique = false
			}
		}

		if isUnique {
			uniqueValues = append(uniqueValues, arr[i])
		}
	}

	return uniqueValues
}

func main() {
	arr := [...]int{1, 2, 3, 1, 2, 3, 4, 5}

	arr1 := printUniqueValues(arr)
	fmt.Println("The unique values :- ", arr1)
}
