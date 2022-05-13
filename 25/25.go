package _5

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{
		Val:  -1,
		Next: nil,
	}
	temp := &head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			temp = temp.Next
			l1 = l1.Next
		} else {
			temp.Next = l2
			temp = temp.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return head.Next
}
