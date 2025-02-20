package main

import (
	"fmt"
)

// Binary (base 2) number system to Decimal (base 10) number system.
func dectoBinary(decimal int) int {
	var ans int
	pow := 1 // 10 ^ 0
	var remainder int

	for decimal > 0 {
		remainder = (decimal % 2)
		decimal = (decimal / 2) // Next Quotient

		ans += (remainder * pow)
		pow *= 10
	}

	return ans
}

func main() {
	fmt.Println("Enter the Decimal Number System to convert into the Binary Number System")
	var num int

	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println(err)
		return
	}

	answer := dectoBinary(num)
	fmt.Println("The Binary Number of given output is :- ", answer)
}
