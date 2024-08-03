package main

import (
	"errors"
	"fmt"
)

func main() {

	numerator, denominator := 31, 0
	var result, remainder, err = intDivision(numerator, denominator)
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("Division result: %v", result)
	default:
		fmt.Printf("Division result: %v\nDivision remainder: %v", result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
