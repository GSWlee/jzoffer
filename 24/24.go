package _4

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre, temp, next := head, head, head
	pre = nil
	next = next.Next
	for temp != nil {
		temp.Next = pre
		pre = temp
		temp = next
		if next != nil {
			next = next.Next
		}
	}
	return pre
}
