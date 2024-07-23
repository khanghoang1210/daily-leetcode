package main

import "fmt"

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
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
	nums := []int{3, 2, 2, 3}
	val := 3
	// Running test cases

	fmt.Printf("Input: nums = %d\n", nums)
	result := removeElement(nums, val)

	fmt.Printf("Output: %v\n", result)

	fmt.Println()

}
