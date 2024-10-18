package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertToASCII(inp string) {

	fmt.Println("Using Range Loop to Convert into ASCII code")
	for _, v := range inp {
		// append(runeType, v)
		fmt.Printf("%c :- %d, ", v, v)
	}
	fmt.Println()

	fmt.Println("Using For Loop to Convert into ASCII code")
	for i := 0; i < len(inp); i++ {
		fmt.Printf("%c : %v, ", inp[i], inp[i])
	}
	fmt.Println()

	fmt.Println("Using []byte to Convert into ASCII code")
	asciiCodes := []byte(inp)
	fmt.Println(asciiCodes)

}

func main() {
	fmt.Println("Enter the string that you want to convert in ASCII code :- ")
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')

	trimInput := strings.TrimSpace(inputString)

	convertToASCII(trimInput)

}

// Format 	Specifier							Meaning	Example
// %c		Character(from ASCII/Unicode)		fmt.Printf("%c", 65) → A
// %d		Decimal integer						fmt.Printf("%d", 123) → 123
// %v		Default format (any type)			fmt.Printf("%v", 123) → 123, fmt.Printf("%v", "Hello") → Hello
