package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return " "
	}
	base := strs[0]
	for i := 0; i < len(base); i++ {
		for _, v := range strs[1:] {
			if i == len(v) || v[i] != base[i] {
				return base[0:i]
			}
		}

	}
	return base
}

func main() {
	// Test cases
	strs := []string{"flower", "flow", "flight"}
	// Running test cases

	fmt.Printf("Input: strs = %s\n", strs)
	result := longestCommonPrefix(strs)

	fmt.Printf("Output: %v\n", result)

	fmt.Println()

}