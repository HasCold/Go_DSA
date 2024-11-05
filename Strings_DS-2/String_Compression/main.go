// Example 1:

// Input: chars = ["a","a","b","b","c","c","c"]
// Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
// Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".

// Return the size of length

package main

import (
	"fmt"
	"strconv"
)

func compress(char []byte) []byte {
	var read = char[0] // Present character
	write := 0         // Frequency for character
	idx := 0           // Present Position
	trackIdx := 0

	if len(char) == 1 {
		return char
	}

	for i := 0; i < len(char); i++ {
		trackIdx = i
		if read == char[i] {
			write++

		} else if read != char[i] {
			if write > 1 {
				if write >= 10 {
					str := strconv.Itoa(write)
					typeCast := []byte(str)

					for i := 1; i <= len(typeCast); i++ {
						char[idx+i] = typeCast[i-1]
					}
				}
				// {'a', 'a', 'a', 'b', 'b', 'a', 'a'}
				str := strconv.Itoa(write)
				typeCast := []byte(str)
				for _, s := range typeCast {
					char[idx] = read
					char[idx+1] = s
				}
			} else {
				char[idx] = read
			}

			if write > 1 {
				// Now updating to the next character
				read = char[i]
				idx = idx + 2
				write = 1
			} else {
				// Now updating to the next character
				read = char[i]
				idx = i
				write = 1
			}

		}

	}

	if trackIdx == len(char)-1 {
		if write >= 10 {
			str := strconv.Itoa(write)
			typeCast := []byte(str)
			// fmt.Println(typeCast)

			for i := 1; i <= len(typeCast); i++ {
				char[idx+i] = typeCast[i-1]
			}
			return char[:idx+len(typeCast)+1] // return char from 0 to idx+len(typeCast) index

		} else if write > 1 {
			str := strconv.Itoa(write)
			typeCast := []byte(str)
			for _, s := range typeCast {
				char[idx] = read
				char[idx+1] = s
			}
			return char[:idx+2]

		} else if write == 1 {
			char[idx] = read
			return char[:idx+1]
		}
	}

	return char
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	chars2 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	chars3 := []byte{'a'}
	chars4 := []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}
	chars5 := []byte{'a', 'a', 'a', 'a', 'a', 'b'}

	out1 := compress(chars)
	out2 := compress(chars2)
	out3 := compress(chars3)
	out4 := compress(chars4)
	out5 := compress(chars5)

	fmt.Println("----------------------------", string(out1))
	fmt.Println("----------------------------", string(out2))
	fmt.Println("----------------------------", string(out3))
	fmt.Println("----------------------------", string(out4))
	fmt.Println("----------------------------", string(out5))
}
