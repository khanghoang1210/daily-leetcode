package main


import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {

    if root == nil {
        return false
    }
    targetSum -= root.Val
    if root.Right == nil && root.Left == nil {
        return targetSum == 0
    }
    return hasPathSum(root.Left,targetSum) || hasPathSum(root.Right, targetSum)
}   

func main(){
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}

	result := hasPathSum(root,22)

	// Print the result
	fmt.Println(result) // Output: true
}