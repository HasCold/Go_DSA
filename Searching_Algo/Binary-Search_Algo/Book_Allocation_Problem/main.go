// --------------------------- Hard Level Problem -----------------------------

// You have n books, each with arr[i] a number of pages. m students need to be allocated contiguous books, with each student getting at least one book.
// Out of all the permutations, the goal is to find the permutation where the sum of the maximum number of pages in a book allotted to a student should be the minimum, out of all possible permutations.

// Note: Return -1 if a valid assignment is not possible, and allotment should be in contiguous order (see the explanation for better understanding).

// Examples:

// Input: n = 4, arr[] = [12, 34, 67, 90], m = 2
// Output: 113
// Explanation: Allocation can be done in following ways:
// {12} and {34, 67, 90} Maximum Pages = 191
// {12, 34} and {67, 90} Maximum Pages = 157
// {12, 34, 67} and {90} Maximum Pages =113.
// Therefore, the minimum of these cases is 113, which is selected as the output.
// Input: n = 3, arr[] = [15, 17, 20], m = 5
// Output: -1
// Explanation: Allocation can not be done.
// Expected Time Complexity: O(n logn)
// Expected Auxilliary Space: O(1)
// Constraints:
// 1 <=  n, m <= 10^5
// 1 <= arr[i] <= 10^6

// --------------------------------------- Raw Explanation ----------------------------------------------------
// arr = [2,1,3,4] , N = 4 and M = 2 (no. of students)
// 1. there are 4 books in the arr and each book has correspondance no. of pages in the order of index
// 2. Each student has to be allocated at least one book.
// 3. m students need to be allocated contiguous books / allotment should be in contigous order :-
// Permuatations :-
// 4. S1 = 2 , S2 = 8   ;  maxNoPages = 8
// 	  S1 = 3 , S2 = 7	;  maxNoPages = 7
// 	  S1 = 6 , S2 = 4	;  maxNoPages = 6  (So this permutation is our answer)
// 5. Out of all the permutations, the goal is to find the permutation where the sum of the maximum number of pages in a book allotted to a student should be the minimum, out of all possible permutations.

// Condition at which the algorithm possibly fails suppose we have 4 books in array arr = [2,1,3,4]  and we have 5 no. of students M=5
// So we cannot allocate 4 books to the 5 students in a contigous allocation.

package main

import "fmt"

func maximumAllocationOfBooks(arr []int, N int, M int) int { // Time Complexity O(log n)
	// find minimum possible maximum pages
	// Possible Sol :-
	// Each student have possible this condition and our answer must lie in the range between of these :-
	// min = 0 <----> max = sum(allPages)
	// 0,1,2,3,4,5,6,7,8,9,10  (Sum of all pages from the given array)
	// So we can apply the binary search on the above sorted array of possible answers but the 0 and 10 are not possibly going to our answer.

	if M <= len(arr) {

		st := 0
		end := sum(arr) // Total no. of maximum answer
		ans := 0

		for st <= end {
			// To prevent the overflow of integer by the large size indexes
			// Optimize the code we use this formula
			mid := st + (end-st)/2 // No. of pages we can allocate to the student ; Max Allowed Pages

			// Check mid valid or not valid
			if isValid(mid, arr, M) { // if true then we find out at the right-search space
				ans = mid
				end = mid - 1
			} else { // if false then we find out at the left-search space
				st = mid + 1
			}

		}
		return ans
	}

	return -1
}

// 0,1,2,3,4,5,6,7,8,9,10 ==> Total sum of array
func isValid(maxAllowedPages int, arr []int, M int) bool { // constant space complexity = O(logN + n)
	// [2, 1, 3, 4]
	// M  --->> No. of students should the book allocate
	// mid  --->> Max Allowed Pages
	stud := 1 // No. of student initialize which is going to be equal M
	pages := 0

	for i := 0; i < len(arr); i++ { // O(log n)

		// Edge Case
		if arr[i] > maxAllowedPages { // Check for the individual book that doesn't have the page greater than maxAllowedPages
			return false
		}
		// Edge Case

		if pages+arr[i] <= maxAllowedPages {
			pages += arr[i]

		} else {
			stud++
			pages = arr[i]
		}

	}
	if stud <= M {
		return true
	} else {
		return false
	}
}

func sum(arr []int) int { // O(n)
	sum := 0
	for _, val := range arr {
		sum += val
	}

	return sum
}

func main() {

	arr := []int{2, 1, 3, 4} // Books with a respective no. of pages at their index
	N := len(arr)            // No of books
	M := 2                   // No. of students

	fmt.Println("The maximum no. of book pages allocated at a minimum permutation :-", maximumAllocationOfBooks(arr, N, M))
}
