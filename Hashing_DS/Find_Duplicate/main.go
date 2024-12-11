package main

import "fmt"

// T.C := O(n)
// S.C := O(n)  because of the use of Hash Table

func findDuplicate(arr []int) int {

	find := map[int]int{} // key, value pair
	var a = -1

	for i := 0; i < len(arr); i++ {
		num := arr[i]
		_, exist := find[num]

		if exist {
			a = num
			return a
		}
		find[num] = i
	}

	return a

}

func main() {
	arr := []int{3, 1, 3, 4, 2}

	fmt.Println(findDuplicate(arr))
}
