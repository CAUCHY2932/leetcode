package leetcode213

type ListNode struct {
	Val  int
	Next *ListNode
}

// dummy node
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}
	fast, slow := dummyNode, dummyNode
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyNode.Next
}
