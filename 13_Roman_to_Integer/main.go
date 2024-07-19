package main

import "fmt"

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i == 0 {
			result = m[s[i]] + result
		} else if m[s[i]] > m[s[i-1]] {
			result = (m[s[i]] - m[s[i-1]]) + result
			i--
		} else {
			result = m[s[i]] + result
		}

	}
	return result
}

func main() {
	// Test cases
	s := "MCMXCIV"
	// Running test cases

	fmt.Printf("Input: x = %s\n", s)
	result := romanToInt(s)

	fmt.Printf("Output: %v\n", result)

	fmt.Println()

}
