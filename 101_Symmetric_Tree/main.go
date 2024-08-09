package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}
func isSymmetric(root *TreeNode) bool {
   
    if root.Left == nil && root.Right == nil{
        return true
    }
    q := root.Right
    p := root.Left
    return isSymmetricTree(p,q)

}
func isSymmetricTree(p *TreeNode, q * TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if (p == nil && q != nil) || (p != nil && q == nil) {
        return false
    }
    if p.Val != q.Val {
        return false
    }

    return isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}

	result := isSymmetric(root)

	// Print the result
	fmt.Println(result) // Output: true
}