package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func threeSumBruteApproach(arr []int) [][]int {

	vec := [][]int{}
	seen := map[string]bool{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {

				if arr[i]+arr[j]+arr[k] == 0 {

					triplet := []int{arr[i], arr[j], arr[k]}
					sort.Ints(triplet)
					key := strings.Join(strings.Fields(fmt.Sprint(triplet)), ",")

					if !seen[key] {
						vec = append(vec, triplet) // to dynamically add a new triplet to the slice.
						seen[key] = true
					}
				}
			}
		}
	}

	return vec
}

// Better Approach
// Intuition :-  a + b + c = target
// if we pick first => a then b + c = -a = target

// HashTable Implementation (Set_DS)
// In Map and Set, there is a hashing data structure is used

// type OrderedSet struct {
// 	set map[string]struct{}
// }

func betterApproachThreeSum(arr []int) [][]int { // {-1, 0, 1, 2, -1, -4}

	uniqueTriplet := map[string]bool{}
	ans := [][]int{}

	for i := 0; i < len(arr); i++ {
		target := -arr[i]
		s := map[string]bool{}

		for j := i + 1; j < len(arr); j++ {
			third := target - arr[j]
			thrd := strconv.Itoa(third)
			_, exist := s[thrd]

			if exist {
				triplet := []int{arr[i], arr[j], third}
				sort.Ints(triplet)
				key := strings.Join(strings.Fields(fmt.Sprint(triplet)), ",")

				if !uniqueTriplet[key] {
					uniqueTriplet[key] = true
					ans = append(ans, triplet)
				}

			}
			s[strconv.Itoa(arr[j])] = true
		}
	}
	return ans
}

// Intuition => Sort the array
// i, j, k
func twoPointerApproach(arr []int) [][]int { // time Complexity => O(n logn) ; space complexity => O(unique Triplet)

	sort.Ints(arr) // [-4, -1, -1, 0, 1, 2]
	ans := [][]int{}
	n := len(arr)

	for i := 0; i < len(arr); i++ {
		// To avoid the duplicate of triplet
		if i > 0 && arr[i] == arr[i-1] {
			continue //continue keyword in Go is used to skip the current iteration of a loop and move to the next iteration.
		}

		j := i + 1
		k := n - 1 // end element of array

		for j < k {
			num := arr[i] + arr[j] + arr[k]

			if num == 0 {
				ans = append(ans, []int{arr[i], arr[j], arr[k]})
				j++
				k--

				// To avoid the duplicate of triplet
				for j < k && arr[j] == arr[j-1] {
					j++
				}

			}

			if num < 0 {
				j++
			} else if num > 0 {
				k--
			}
		}
	}

	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println("Brute Force Approach :- ", threeSumBruteApproach(nums))

	fmt.Println("Better Approach :- ", betterApproachThreeSum(nums))

	fmt.Println("Two-Pointer Optimized Approach :- ", twoPointerApproach(nums))
}
