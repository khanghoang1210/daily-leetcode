package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
 type ListNode struct {
	Val int
	Next *ListNode
}
 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil{
        return list2
    }else if list2 == nil {
        return list1
    }
    current := list1
    for current.Next != nil {
        current = current.Next
    }
    current.Next = list2

    for i := list1; i != nil; i = i.Next {
        for j := i.Next; j != nil; j = j.Next{
            if i.Val > j.Val {
                swap(i,j)
            }
        }
    }

    return list1
}
func swap(a, b *ListNode) {
    temp := a.Val
    a.Val = b.Val
    b.Val = temp
}
func printList(head *ListNode) {
    for head != nil {
        fmt.Printf("%d -> ", head.Val)
        head = head.Next
    }
    fmt.Println("nil")
}
func main() {
    // Create an unsorted linked list
    head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 4}

	head1 := &ListNode{Val: 1}
    head1.Next = &ListNode{Val: 3}
    head1.Next.Next = &ListNode{Val: 4}


    fmt.Println("Unsorted list:")
    printList(head)
	printList(head1)

    // Sort the linked list
    head = mergeTwoLists(head, head1)

    fmt.Println("Sorted list:")
    printList(head)
}
