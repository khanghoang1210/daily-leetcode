package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
func sortedArrayToBST(nums []int) *TreeNode {
    root := &TreeNode{}
    if len(nums) == 0 {
        return nil
    }
    mid := len(nums)/2
    root.Val = nums[mid]
    if mid < 0 {
        return root
    }
    root.Left = sortedArrayToBST(nums[:mid])
    root.Right = sortedArrayToBST(nums[mid+1:])
    return root
}

// Function to print the tree using preorder traversal (Root, Left, Right)
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}

	fmt.Printf("Input: nums = %v\n", nums)
	result := sortedArrayToBST(nums)

	fmt.Printf("Output (Preorder Traversal): ")
	preorderTraversal(result)
	fmt.Println()
}