package leetcode

func removeZeroSumSublists(head *ListNode) *ListNode {
    dummy := &ListNode{0, head}
    sum := 0
    m := make(map[int]*ListNode)
    for i := dummy; i != nil; i = i.Next {
         sum += i.Val
         m[sum] = i
    }
    sum = 0
    for i := dummy; i != nil; i = i.Next {
        sum += i.Val
        i.Next = m[sum].Next
    }
    return dummy.Next
}
