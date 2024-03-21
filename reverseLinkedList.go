package leetcode

func reverseList(head *ListNode) *ListNode {
	var previous *ListNode

	for head != nil {
		next := head.Next
		head.Next = previous
		previous = head
		head = next
	}

	return previous
}
