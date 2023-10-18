package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getGreatestCommonDivisor(a, b int) int {
	if a < b {
		a, b = b, a
	}

	for b > 0 {
		a, b = b, a%b
	}

	return a
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	orig := head.Next

	for orig != nil {
		gcd := getGreatestCommonDivisor(cur.Val, orig.Val)
		newNode := &ListNode{Val: gcd, Next: orig}
		cur.Next = newNode
		cur = orig
		orig = cur.Next
	}

	return head
}
