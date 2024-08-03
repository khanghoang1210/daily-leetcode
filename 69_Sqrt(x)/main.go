package main

import "fmt"

func mySqrt(x int) int {
	left := 0
	right := x - 1
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x || (mid*mid < x && (mid+1)*(mid+1) > x) {
			return mid
		}
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	x := 10000000
	// Running test cases

	fmt.Printf("Input: x = %d\n", x)
	result := mySqrt(x)

	fmt.Printf("Output: %v\n", result)

	fmt.Println()

}
