// This is a practice column for a Arrays_DS-1 where you can test your skills

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the base (x)")
	reader := bufio.NewReader(os.Stdin)

	inp, _ := reader.ReadString('\n')
	trimInp := strings.TrimSpace(inp)
	Base := strings.Map(func(r rune) rune {
		// to check whther we get the right decimal place
		if r >= '0' && r <= '9' {
			return r
		}

		return -1
	}, trimInp)
	baseInt, err := strconv.ParseFloat(Base, 64)
	if err != nil {
		fmt.Println("Invalid Input for base")
		return
	}

	fmt.Println("Enter the exponent (n)")
	expInp, _ := reader.ReadString('\n')
	trimExpInp := strings.TrimSpace(expInp)
	Exponent := strings.Map(func(r rune) rune {
		// rune data-type will return the ASCII code means mathematical representation of any character
		if r >= '0' && r <= '9' {
			return r
		}

		return -1
	}, trimExpInp)
	exponentInt, err := strconv.Atoi(Exponent)
	if err != nil {
		fmt.Println("Invalid Input for exponent")
		return
	}

	fmt.Printf("The binary exponentiation of base %v with exponent %v is %v", baseInt, exponentInt, pow(baseInt, exponentInt))
}

func pow(base float64, exponent int) float64 {
	// We will run the loop on exponent of binary form
	// binaryForm := decToBinary(exponent)
	var ans = 1.0
	binaryForm := int(exponent)

	for binaryForm > 0 {
		if binaryForm%2 == 1 {
			ans *= base
		}

		base *= base
		binaryForm = (binaryForm / 2)
	}

	return ans
}
