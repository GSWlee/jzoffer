package _2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	now := []*TreeNode{root}
	for len(now) > 0 {
		temp := now
		now = []*TreeNode{}
		t := []int{}
		for _, v := range temp {
			t = append(t, v.Val)
			if v.Left != nil {
				now = append(now, v.Left)
			}
			if v.Right != nil {
				now = append(now, v.Right)
			}
		}
		ans = append(ans, t)
	}
	return ans
}
