package _4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	temp := []int{}
	ans := [][]int{}
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		temp = append(temp, root.Val)
		sum += root.Val
		if sum == target && root.Right == nil && root.Left == nil {
			ans = append(ans, append([]int{}, temp...))
		}
		dfs(root.Left, sum)
		dfs(root.Right, sum)
		sum -= root.Val
		temp = temp[:len(temp)-1]
	}
	dfs(root, 0)
	return ans
}
