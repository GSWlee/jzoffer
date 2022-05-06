package _7

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
//  buildTree
//  @Description: 利用前序遍历与中序遍历来生成原始二叉树
//  @param preorder： 前序遍历
//  @param inorder： 中序遍历
//  @return *TreeNode：
//
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else {
		temp := TreeNode{
			Val:   preorder[0],
			Left:  nil,
			Right: nil,
		}
		t := -1
		for k, v := range inorder {
			if v == preorder[0] {
				t = k
				break
			}
		}
		temp.Left = buildTree(preorder[1:1+t], inorder[:t])
		temp.Right = buildTree(preorder[1+t:], inorder[t+1:])
		return &temp
	}
}
