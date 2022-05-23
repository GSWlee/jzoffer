package _2

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ptra, ptrb := headA, headB
	if headA == nil {
		return nil
	}
	if headB == nil {
		return nil
	}
	flag1 := 0
	flag2 := 0
	for ptra != ptrb {
		if ptra == nil {
			if flag1 == 0 {
				ptra = headB
			}
		} else {
			ptra = ptra.Next
		}
		if ptrb == nil {
			if flag2 == 0 {
				ptrb = headA
			}
		} else {
			ptrb = ptrb.Next
		}
	}
	return ptra
}
