package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Convert 101 binary number (Base 2) to decimal number (Base 10) :-

// 10 | 101  = 1
// 10 | 10  = 0
// 10 | 1  = 1
// 	 	0

// 1 * 2^2(4) + 0 * 2^1(2) + 1 * 2^0(1) = 5

func binaryToDec(binaryNum int) {

	pow := 1 // 2^0(1) if * 2 = 2^1(2) if * 2 = 2^2(4)
	var rem int
	var ans int

	for binaryNum > 0 {
		rem = binaryNum % 10
		binaryNum = binaryNum / 10 // quotient

		ans += rem * pow
		pow *= 2
	}

	fmt.Printf("The Decimal Number System After Conversion is %v \n", ans)
}

func isVowel(userInput string) {
	var vowels = "aeiouAEIOU"
	for _, inp := range userInput {
		if strings.Contains(vowels, string(inp)) {
			fmt.Println("The vowels contain :- ", string(inp))
		}
	}
}

func main() {
	fmt.Println("Enter the Binary Number (only 0s and 1s) for the conversion in Decimal Number System :- ")

	// IO bound operations
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	trimInput := strings.TrimSpace(input)

	isVowel(input)

	// strings.Map() is a powerful function that applies a given function (mapper) to each character (rune) in a string and returns a new string with the transformed characters.
	binaryStr := strings.Map(func(r rune) rune {
		if r == '0' || r == '1' {
			return r
		}
		return -1 // Remove other characters by returning -1
	}, trimInput)

	// fmt.Println("The nm :-", input[0]) // The character '7' has an ASCII value of 55.
	// So when you access input[0], it is returning the byte representing the character '7', which is 55 in ASCII.
	// 1 byte = 8 bits unsigned integer (uint8) containing 256 values , range from 0 to 255. Bytes are often used to represent characters, binary data, or raw memory.

	// uint8 = 0 to 255 values
	// int8 =  (-128) to (127) values

	fmt.Println("After removing unnecessary characters :- ", binaryStr)
	convertStr, _ := strconv.Atoi(binaryStr)
	binaryToDec(convertStr)
}
