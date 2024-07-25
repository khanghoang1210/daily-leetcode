package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	index := 0
	isLeft := true
	res := 0
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
			index = mid
			isLeft = false
		} else {
			right = mid - 1
			index = mid
			isLeft = true
		}
	}
	if isLeft {
		res = index
	} else {
		res = index + 1
	}
	if res < 0 {
		res = 0
	}
	return res
}

func main() {
	// Test cases
	nums := []int{1,3,5,6}
	target := 0
	// Running test cases

	fmt.Printf("Input: nums = %d\n", nums)
	fmt.Printf("Input: target = %d\n", target)
	result := searchInsert(nums, target)

	fmt.Printf("Output: %v\n", result)

	fmt.Println()

}