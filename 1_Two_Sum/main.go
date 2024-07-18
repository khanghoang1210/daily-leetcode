package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i)
				res = append(res, j)
			}
		}
	}
	return res

}

func main() {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
		{[]int{1, 2, 3, 4, 5}, 9},
	}

	// Running test cases
	for _, tc := range testCases {
		fmt.Printf("Input: nums = %v, target = %d\n", tc.nums, tc.target)
		result := twoSum(tc.nums, tc.target)
		if len(result) == 2 {
			fmt.Printf("Output: %v\n", result)
		} else {
			fmt.Println("No solution found")
		}
		fmt.Println()
	}
}
