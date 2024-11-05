// Given an input string s, reverse the order of the words.

// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

// Return a string of the words in reverse order concatenated by a single space.

// Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

// Example 1:

// Input: s = "the sky is blue"
// Output: "blue is sky the"

package main

import (
	"fmt"
	"strings"
)

func reversedWordsInStrings(s string) string {
	trimInput := strings.TrimSpace(s)

	fieldInput := strings.Fields(trimInput)

	// Two-pointer Approach
	i := 0
	j := len(fieldInput) - 1

	for i <= j {
		fieldInput[i], fieldInput[j] = fieldInput[j], fieldInput[i]
		i++
		j--
	}

	return strings.Join(fieldInput, " ")
}

func main() {
	s := "the sky is blue"
	s2 := "  hello world  "

	fmt.Println("The reversed word of string :- ", reversedWordsInStrings(s))
	fmt.Println("The reversed word of string :- ", reversedWordsInStrings(s2))
}
