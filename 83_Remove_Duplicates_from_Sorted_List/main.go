package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
  }

 func deleteDuplicates(head *ListNode) *ListNode {
    current := head
    if head == nil {
        return head
    }
    if current.Next == nil {
        return head
    }

    for current != nil && current.Next != nil {
        if current.Val != current.Next.Val {
            current = current.Next
        }else {
            current.Next = current.Next.Next
        }
        
    }
    return head
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
    head.Next.Next.Next = &ListNode{Val: 4}


    fmt.Println("Duplicates Sorted list:")
    printList(head)

    // Sort the linked list
    head = deleteDuplicates(head)

    fmt.Println("Sorted list after remove duplicates:")
    printList(head)
}