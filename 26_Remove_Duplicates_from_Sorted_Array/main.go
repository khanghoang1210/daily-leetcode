package main

import "fmt"

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = remove(nums, i)
			i--
		}
	}

	return len(nums)
}
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	// Test cases
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// Running test cases

	fmt.Printf("Input: nums = %d\n", nums)
	result := removeDuplicates(nums)

	fmt.Printf("Output: %v\n", result)

	fmt.Println()

}
