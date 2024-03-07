package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    if head.Next == nil {
        return head
    }

    if head.Next.Next == nil {
        return head.Next
    }
 
    var slow, fast *ListNode = head, head
    for {
        slow = slow.Next
        fast = fast.Next.Next

        if (fast.Next != nil) && (fast.Next.Next == nil) {
            slow = slow.Next
            break
        }

        if (fast.Next == nil) {
            break
        }
    }

    return slow
}
