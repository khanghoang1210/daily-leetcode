package main

import "fmt"

// accept 67/83 test cases
// func strStr(haystack string, needle string) int {
//     index := 0
//     for i := 0; i <= len(haystack)-1; i++ {
//         if haystack[i] == needle[index]{
//             if index >= len(needle)-1 {
//                 return i - index

//             } else{
//                 index++
//             }

//         } else {
//             index = 0
//             i = i
//         }

//     }
//     return -1
// }

func strStr(haystack string, needle string) int {
    for i := 0; i < len(haystack); i++ {
        if len(needle) + i > len(haystack) {
            return -1
        } else if haystack[i:len(needle)+i] == needle {
            return i
        }
    }
    return -1
}

func main() {
	// Test cases
	haystack := "mississippi"
	needle := "issip"
	// Running test cases

	fmt.Printf("Input: haystack = %s\n", haystack)
	fmt.Printf("Input: needle = %s\n", needle)
	result := strStr(haystack, needle)

	fmt.Printf("Output: %v\n", result)

	fmt.Println()

}
