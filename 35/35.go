package _5

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		} else {
			node.Next.Random = nil
		}
	}
	headNew := head.Next
	tmp := head
	for temp := headNew; temp != nil; temp = temp.Next {
		tmp.Next = temp.Next
		if temp.Next == nil {
			break
		}
		tmp = tmp.Next
		temp.Next = temp.Next.Next
	}
	return headNew
}
