package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	diff := getHeight(root.Right) - getHeight(root.Left)

	if math.Abs(float64(diff)) > 1.0 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(root *TreeNode) float64 {
	if root == nil {
		return 0.0
	}
	return math.Max(getHeight(root.Left), getHeight(root.Right)) + 1.0
}

func main() {
	// Creating a balanced tree
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Left = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 5}
	root1.Right.Left = &TreeNode{Val: 6}
	root1.Right.Right = &TreeNode{Val: 7}

	// Creating an unbalanced tree
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Left.Left = &TreeNode{Val: 3}
	root2.Left.Left.Left = &TreeNode{Val: 4}

	fmt.Println("Is the first tree balanced?", isBalanced(root1)) // Expected output: true
	fmt.Println("Is the second tree balanced?", isBalanced(root2)) // Expected output: false
}
