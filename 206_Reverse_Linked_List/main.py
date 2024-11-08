# Time: O(n) / Space: O(1)

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        prev = None
        current = head
        while current is not None:
            tmp = current.next
            current.next = prev
            prev = current
            current = tmp
        return prev

# Helper function to print a linked list
def print_list(head: ListNode):
    current = head
    while current is not None:
        print(current.val, end=" -> " if current.next else "")
        current = current.next
    print()

# Helper function to create a linked list from a list of values
def create_linked_list(values):
    if not values:
        return None
    head = ListNode(values[0])
    current = head
    for value in values[1:]:
        current.next = ListNode(value)
        current = current.next
    return head

def main():
    # Create a linked list from a list of values
    values = [1, 2, 3, 4, 5]
    head = create_linked_list(values)
    
    print("Original list:")
    print_list(head)

    # Reverse the linked list
    solution = Solution()
    reversed_head = solution.reverseList(head)
    
    print("Reversed list:")
    print_list(reversed_head)

if __name__ == "__main__":
    main()
