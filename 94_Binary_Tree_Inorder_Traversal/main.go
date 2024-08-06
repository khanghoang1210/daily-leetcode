package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack []*TreeNode

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	s := make(stack, 0)

	current := root
	for current != nil || len(s) != 0 {
		for current != nil {
			s.Push(current)
			current = current.Left
		}
		current, _ = s.Pop()
		res = append(res, current.Val)
		current = current.Right
	}
	return res
}

func (s *stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *stack) Pop() (*TreeNode, bool) {
	if len(*s) == 0 {
		return nil, false
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top, true
}

func main() {
	// Create a sample binary tree
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	// Perform inorder traversal
	result := inorderTraversal(root)

	// Print the result
	fmt.Println(result) // Output: [4 2 5 1 3]
}
