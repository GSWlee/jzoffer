package _4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	ans := -1
	flag := false
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil || flag {
			return
		}
		dfs(root.Right)
		k--
		if k == 0 {
			ans = root.Val
			flag = true
			return
		}
		dfs(root.Left)
	}
	dfs(root)
	return ans
}
