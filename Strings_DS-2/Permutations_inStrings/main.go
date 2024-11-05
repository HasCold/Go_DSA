package main

import "fmt"

// Permutations :- No. of arrangements OR no. of orders

// Given two strings s1 and s2, return true if s2 contains a
// permutation
//  of s1, or false otherwise.

// In other words, return true if one of s1's permutations is the substring of s2.

// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").

func isFreqSame(freq, windFreq [26]int) bool { // O(1)

	for i := 0; i < len(windFreq); i++ {
		if freq[i] != windFreq[i] {
			return false
		}
	}

	return true
}

func permutationsInStrings(s1, s2 string) bool {
	freq := [26]int{0}

	for i := 0; i < len(s1); i++ {
		// idx := s1[i] - 'a'   //  'a'-'a' = 0 , 'b'-'a' = 1 ....
		// idx++
		freq[s1[i]-'a']++
	}

	var windSize = len(s1) // 2

	for i := 0; i < len(s2); i++ {
		windIdx := 0           // window Index specifically for two characters
		idx := i               // String Index
		windFreq := [26]int{0} // you are creating an array with 26 integer elements, all initialized to 0 by default.

		for windIdx < windSize && idx < len(s2) {
			position := s2[idx] - 'a'
			windFreq[position]++

			windIdx++
			idx++
		}

		if isFreqSame(freq, windFreq) {
			return true
		}
	}

	return false

}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"

	fmt.Println("The Permutations in Strings :- ", permutationsInStrings(s1, s2))
}
