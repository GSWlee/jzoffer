package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	ans, temp := []int{}, []int{}
	ptr := head
	for ptr != nil {
		temp = append(temp, ptr.Val)
		ptr = ptr.Next
	}
	for j := len(temp) - 1; j >= 0; j-- {
		ans = append(ans, temp[j])
	}
	return ans
}
