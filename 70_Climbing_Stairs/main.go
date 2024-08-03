package main

import "fmt"


func climbStairs(n int) int {
    one := 1
    two := 1
    total := 0
    if n == 1 {
        return 1
    }
    for i:= 2; i <= n; i++ {
       total = one + two
       two = one
       one = total
    }
   
    return total
}

func main() {
	n := 45
	// Running test cases

	fmt.Printf("Input: n = %d\n", n)
	result := climbStairs(n)

	fmt.Printf("Output: %v\n", result)

	fmt.Println()
}