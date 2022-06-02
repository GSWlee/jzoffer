package _8

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	roadp, roadq, temp := []*TreeNode{}, []*TreeNode{}, []*TreeNode{}
	var dfs func(root, q *TreeNode) bool
	dfs = func(root, q *TreeNode) bool {
		if root == nil {
			return false
		}
		if root.Val == q.Val {
			temp = append(temp, root)
			return true
		} else {
			temp = append(temp, root)
			flag := dfs(root.Left, q)
			if flag {
				return flag
			}
			flag = dfs(root.Right, q)
			if flag {
				return flag
			}
			temp = temp[:len(temp)-1]
			return false
		}
	}
	dfs(root, p)
	roadp = temp
	temp = []*TreeNode{}
	dfs(root, q)
	roadq = temp

	for i := 0; true; i++ {
		if i+1 < len(roadp) && i+1 < len(roadq) && roadp[i+1].Val == roadq[i+1].Val {
			continue
		} else {
			return roadp[i]
		}
	}
	return root
}
