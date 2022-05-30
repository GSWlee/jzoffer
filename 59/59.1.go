package _9

func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	ans := []int{}
	stack := [][]int{}
	l := 0
	for i := 0; i < len(nums); i++ {
		if l < k {
			for len(stack) > 0 {
				if stack[len(stack)-1][0] < nums[i] {
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			if len(stack) == 0 || stack[len(stack)-1][0] >= nums[i] {
				stack = append(stack, []int{nums[i], i})
			}
			l++
		} else {
			ans = append(ans, stack[0][0])
			if stack[0][1]+k <= i {
				stack = stack[1:]
			}
			for len(stack) > 0 {
				if stack[len(stack)-1][0] < nums[i] {
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			if len(stack) == 0 || stack[len(stack)-1][0] >= nums[i] {
				stack = append(stack, []int{nums[i], i})
			}
		}
	}
	ans = append(ans, stack[0][0])
	return ans
}
