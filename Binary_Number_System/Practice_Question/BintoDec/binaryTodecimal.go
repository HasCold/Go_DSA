package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// basically we have to go towards the Decimal Num System so in this case the divisor will be 10
func BintoDec(input int) int {
	divisor := 10
	var remainder int
	var ans int
	pow := 1 // 2 ^ 0 = 1

	for input > 0 {
		remainder = (input % divisor)
		input = (input / divisor) // next quotient

		ans += (remainder * pow)
		pow *= 2
	}

	return ans
}

func main() {
	fmt.Println("Enter the Binary Number System to convert into the Decimal Number System")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	trimInput := strings.TrimSpace(input)

	// Remove the unwanted characters
	binaryInputStr := strings.Map(func(r rune) rune {
		if r == '0' || r == '1' { // rune data-type will represent the ASCII code of any value or ASCII code means the mathematical representation of any characters.
			return r
		}
		return -1 // remove the characters by returning by -1
	}, trimInput)

	if binaryInputStr == "" {
		log.Fatal("Invalid Binary Number")
		return
	}

	binaryInput, err := strconv.Atoi(binaryInputStr)
	if err != nil {
		log.Println("Error while converting the string to integer :- ", err)
		return
	}

	decimalNum := BintoDec(binaryInput)
	fmt.Println("The Decimal Number of the given input :- ", decimalNum)
}
