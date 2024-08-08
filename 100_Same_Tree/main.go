package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}
 
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if q == nil && p == nil {
        return true
    }
    if (p == nil || q == nil) {
        return false
    }
    if p.Val != q.Val {
        return false
    }

    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Left = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 5}

	result := isSameTree(root, root1)

	// Print the result
	fmt.Println(result) // Output: [4 2 5 1 3]
}