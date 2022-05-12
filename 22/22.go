package _2

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}
	first, secound := head, head
	for i := 0; i < k; i++ {
		if first == nil {
			return nil
		}
		first = first.Next
	}
	for first != nil {
		first = first.Next
		secound = secound.Next
	}
	return secound
}
