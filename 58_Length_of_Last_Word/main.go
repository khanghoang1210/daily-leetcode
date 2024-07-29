package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	if len(s) == 1 {
		return 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			length++
			if i == 0 || s[i-1] == ' ' {
				return length
			}
		}
	}
	return length
}

func main() {
	// Test cases
	s := "   fly me   to   the moon  "
	// Running test cases

	fmt.Printf("Input: nums = %s\n", s)
	result := lengthOfLastWord(s)

	fmt.Printf("Output: %v\n", result)

	fmt.Println()

}