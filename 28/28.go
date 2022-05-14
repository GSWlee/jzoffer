package _8

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	right, left := []*TreeNode{}, []*TreeNode{}
	var br func(root *TreeNode)
	var bl func(root *TreeNode)
	br = func(root *TreeNode) {
		right = append(right, root)
		if root == nil {
			return
		} else {
			br(root.Right)
			br(root.Left)
		}
	}
	bl = func(root *TreeNode) {
		left = append(left, root)
		if root == nil {
			return
		} else {
			bl(root.Left)
			bl(root.Right)
		}
	}
	br(root)
	bl(root)

	for k, v := range right {
		if (v == nil && left[k] != nil) || (v != nil && left[k] == nil) {
			return false
		}
		if v != nil && v.Val != left[k].Val {
			return false
		}

	}
	return true
}
