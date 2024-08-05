package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p3 := m + n - 1

	for p2 >= 0 && p1 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p3] = nums1[p1]
			p1--
		} else {
			nums1[p3] = nums2[p2]
			p2--
		}
		p3--
	}
	for p2 >= 0 {
		nums1[p3] = nums2[p2]
		p3--
		p2--
	}

}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	fmt.Printf("Input: nums1 = %d\n", nums1)
	fmt.Printf("Input: m = %d\n", m)
	fmt.Printf("Input: nums2 = %d\n", nums2)
	fmt.Printf("Input: n = %d\n", n)
	merge(nums1, m, nums2, n)

	fmt.Printf("Output: %v\n", nums1)

	fmt.Println()
}
