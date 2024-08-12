package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
func maxDepth(root *TreeNode) int {
    count := 0.0
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }
    if root.Left != nil || root.Right != nil {
        count += 1.0
    }
    count = count + math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))
    return int(count)
}

func main(){
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	result := maxDepth(root)

	// Print the result
	fmt.Println(result) // Output: 3
}