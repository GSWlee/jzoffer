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
	flag := true
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
		if flag {
			ans = append(ans, t)
			flag = false
		} else {
			tmp := []int{}
			for k := len(t) - 1; k >= 0; k-- {
				tmp = append(tmp, t[k])
			}
			ans = append(ans, tmp)
			flag = true
		}

	}
	return ans
}
