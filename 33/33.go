package _3

func verifyPostorder(postorder []int) bool {
	if len(postorder) < 3 {
		return true
	}
	index := -1
	flag := false
	for i := 0; i < len(postorder)-1; i++ {
		if postorder[i] < postorder[len(postorder)-1] && flag == false {
			index = i
		} else if flag == false {
			flag = true
		} else if postorder[i] < postorder[len(postorder)-1] {
			return false
		} else {
			continue
		}
	}
	t, p := verifyPostorder(postorder[:index+1]), verifyPostorder(postorder[index+1:len(postorder)-1])
	return t && p
}
