package main

import "fmt"

// we can use the nested loop to determine the unique int from an array :-

// e.g. :-
// [1  ,	2,3,1,2,3,4,5]
// 1st loop iterate, 2nd loop to match the remaining one's

func uniqueValueFromArr(arr [8]int) []int {
	var uniqueValue []int // Slices

	for i := 0; i < len(arr); i++ {
		isUnique := true
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == arr[j] {
				isUnique = false
			}
		}

		if isUnique {
			uniqueValue = append(uniqueValue, arr[i])
		}
	}

	return uniqueValue
}

func main() {
	arr := [8]int{1, 2, 3, 1, 2, 3, 4, 5}

	uniqueNums := uniqueValueFromArr(arr)
	fmt.Println("The unique slice :- ", uniqueNums)
}
