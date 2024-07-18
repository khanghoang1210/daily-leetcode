package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	
	reverseString := ""
	for _, v := range strX {
		reverseString = string(v) + reverseString
		
	}

	return reverseString == strX
}

func main() {
	// Test cases
	x := -121
	// Running test cases

	fmt.Printf("Input: x = %d\n", x)
	result := isPalindrome(x)

	fmt.Printf("Output: %v\n", result)

	fmt.Println()
	
}
