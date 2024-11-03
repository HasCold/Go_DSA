package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Strings are stored in memory occupy space around 1 byte.

func stringReversal(arr string) string {
	// Two Pointer Approach
	str := []rune(arr) // rune data-type will type cast the string into ASCII code (Mathematical Representation Of any character)
	i := 0
	j := len(str) - 1

	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}

	return string(str)
}

func wordReversedString(str string) string {
	words := strings.Fields(str) // Fields splits the string s around each instance of one or more consecutive white space characters
	i := 0
	j := len(words) - 1

	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	return strings.Join(words, " ")
}

func main() {
	var str = []string{"Hello From HasCold !"} // shows ASCII code means mathematical representaton of any character.

	// fmt.Println("Enter the Charcters :- ")
	// var inp string
	// fmt.Scan(&str)
	// fmt.Println("The input taken from user :- ", str)

	for i, s := range str {
		str[i] = stringReversal(s)
	}
	fmt.Println("The Alphabet Reversed String :- ", str)

	var str2 = "the sky is blue"
	reversedStr := wordReversedString(str2)
	fmt.Println("The Word Reversed String :- ", reversedStr)

	// IO bound operations
	fmt.Println("Enter the Charcters :- ")
	reader := bufio.NewReader(os.Stdin)
	inpStr, _ := reader.ReadString('\n') // ('\n')  -->> Delimiter means when we press for the next line it will stop taking the command strings from us
	fmt.Println("The input taken from user via io bound operations :- ", inpStr)

	fmt.Println(len(str))
}
