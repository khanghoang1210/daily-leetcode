package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
 func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    } 
    left := minDepth(root.Left)
    right := minDepth(root.Right)

    return 1 + min(left,right)

}


func min(a, b int) int {
    if a == 0 {
        return b
    }
    if b == 0 {
        return a
    }
    if a > b {
        return b
    }
    return a
}

func main(){
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	result := minDepth(root)

	// Print the result
	fmt.Println(result) // Output: 2
}