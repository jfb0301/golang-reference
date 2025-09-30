package main

import (
	"fmt"
	"errors"
	"os"
)

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {


	// Test case 1: Normal division
	quotient, remainder, err := divAndRemainder(10, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 3: Quotient = %d, Remainder = %d\n", quotient, remainder)
	}



	// Test case 2: Division with no remainder
	quotient, remainder, err = divAndRemainder(20, 5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("20 ÷ 5: Quotient = %d, Remainder = %d\n", quotient, remainder)
	}



	// Test case 3: Division by zero
	quotient, remainder, err = divAndRemainder(15, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("15 ÷ 0: Quotient = %d, Remainder = %d\n", quotient, remainder)
	}



	// Test case 4: Negative numerator
	quotient, remainder, err = divAndRemainder(-17, 4)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("-17 ÷ 4: Quotient = %d, Remainder = %d\n", quotient, remainder)
	}

	// Test case 5: Simple

	quotient, remainder, err = divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(quotient, remainder)
}