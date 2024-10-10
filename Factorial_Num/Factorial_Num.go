package main

import (
	"errors"
	"fmt"
	"log"
)

func factorial(num int) (res int, err error) {
	if num < 0 {
		return 0, errors.New("Enter the integer greater than 0")
	}

	if num == 0 || num == 1 {
		return 1, nil
	}

	res, er := factorial(num - 1)

	if er != nil {
		log.Fatal(err)
	}

	return (res * num), nil
}

func main() {
	fmt.Println("Enter the number for factorial")
	var n int
	fmt.Scanln(&n)

	res, err := factorial(n)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("The Factorial of %v is %v \n", n, res)
	}
}
