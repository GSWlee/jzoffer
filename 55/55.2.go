package _5

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	ans := true
	flag := false

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil || false {
			return 0
		}
		l, r := maxDepth(root.Left), maxDepth(root.Right)
		if l > r {
			if l-r > 1 {
				ans = false
				flag = true
			}
			return l + 1
		} else {
			if r-l > 1 {
				ans = false
				flag = true
			}
			return r + 1
		}
	}
	maxDepth(root)
	return ans
}
