package _2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	queue := []*TreeNode{}
	ans := []int{}
	if root == nil {
		return ans
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		ans = append(ans, temp.Val)
		if temp.Left != nil {
			queue = append(queue, temp.Left)
		}
		if temp.Right != nil {
			queue = append(queue, temp.Right)
		}
	}
	return ans
}
