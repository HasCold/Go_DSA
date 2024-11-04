package main

import (
	"fmt"
	"strings"
)

// Vald Palindrome String -->> Means when we read the string from forward or from backward it still remains same.
// eg :- race car

// We have to check the alpha-numeric characters in strings like a-z, 0-9
// Special Characters := $,?,!

func isAlphaNumeric(ch byte) bool {

	// We will check the alpha-numeric characters for lowercase char from a to z and 0 to 9
	if (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') {
		return true
	}

	return false
}

func isValidPalindrome(s string) bool {
	lowerCaseStr := strings.ToLower(s)

	i := 0
	j := len(s) - 1

	// a-z , A-Z and 0-9
	for i < j {

		if !isAlphaNumeric(lowerCaseStr[i]) {
			i++
			continue
		} else if !isAlphaNumeric(lowerCaseStr[j]) {
			j--
			continue
		} else if lowerCaseStr[i] != lowerCaseStr[j] {
			return false
		}

		// Move both pointers towards the center
		i++
		j--
	}

	return true
}

func removePart(str string, start, erase int) string {
	if len(str) < 0 {
		return str
	}

	return str[:start] + str[start+erase:]
}

func removeAllOccurrence(s, part string) string {

	occurrrenceStr := s

	// Find the leftmost occurrence part from the given string
	for len(s) > 0 && strings.Index(occurrrenceStr, part) < len(occurrrenceStr) && strings.Index(occurrrenceStr, part) > -1 {
		occurrrenceStr = removePart(occurrrenceStr, strings.Index(occurrrenceStr, part), len(part))
	}

	return occurrrenceStr
}

func main() {
	// s := "Ac3?e3c&a"
	s := "race a car"
	// fmt.Println(s[1])

	fmt.Println("The given valid palindrome is :- ", isValidPalindrome(s))

	//------------------------------------------- Remove All Occurrence -------------------------------------------
	str := "daabcbaabcbc"
	part := "abc"
	removedStr := removeAllOccurrence(str, part)
	fmt.Println("The remove occurrence part from strings :- ", removedStr)

	s2 := "axxxxyyyyb"
	part2 := "xy"
	removedStr = removeAllOccurrence(s2, part2)
	fmt.Println("The remove occurrence part from strings :- ", removedStr)

}
