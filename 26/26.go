package _6

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSub(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val == B.Val {
		return isSub(A.Left, B.Left) && isSub(A.Right, B.Right)
	} else {
		return false
	}
}

func isSubStructur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val == B.Val {
		return (isSub(A.Left, B.Left) && isSub(A.Right, B.Right)) || isSubStructur(A.Left, B) || isSubStructur(A.Right, B)
	} else {
		return isSubStructur(A.Left, B) || isSubStructur(A.Right, B)
	}
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	if A == nil {
		return false
	}
	if A.Val == B.Val {
		return (isSub(A.Left, B.Left) && isSub(A.Right, B.Right)) || isSubStructur(A.Left, B) || isSubStructur(A.Right, B)
	} else {
		return isSubStructur(A.Left, B) || isSubStructur(A.Right, B)
	}
}
