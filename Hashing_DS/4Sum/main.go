package main

import (
	"fmt"
	"sort"
)

// Condition = i != j != k != l
// 1. Sort the array
// first + second + third + fourth = target
func fourSum(arr []int, tar int) [][]int {

	sort.Ints(arr) // {-2, -1, -1, 1, 1, 2, 2}
	n := len(arr)
	ans := [][]int{}

	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		for j := i + 1; j < len(arr); j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}

			k := j + 1
			l := n - 1

			for k < l {
				sum := arr[i] + arr[j] + arr[k] + arr[l]

				if sum == tar {
					ans = append(ans, []int{arr[i], arr[j], arr[k], arr[l]})
					// Keep the two pointers update towards the next possible target
					k++
					l--

					for k < l && arr[k] == arr[k-1] {
						k++
					}
					for k < l && arr[l] == arr[l+1] {
						l--
					}

				} else if sum < tar { // -2 < 0
					k++
				} else if sum > tar { // 2 > 0
					l--
				}
			}
		}

	}

	return ans
}

func main() {
	arr := []int{-2, -1, -1, 1, 1, 2, 2}
	target := 0

	fmt.Println("The optimized approach for 4Sum :- ", fourSum(arr, target))
}
