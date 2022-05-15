package _1

func validateStackSequences(pushed []int, popped []int) bool {
	ptrpush, ptrpop := 0, 0
	stack := []int{}
	for ptrpop < len(popped) {
		if len(stack) == 0 || stack[len(stack)-1] != popped[ptrpop] {
			for ptrpush < len(pushed) {
				stack = append(stack, pushed[ptrpush])
				if pushed[ptrpush] == popped[ptrpop] {
					ptrpush++
					break
				}
				ptrpush++
			}
		}
		if len(stack) == 0 || stack[len(stack)-1] != popped[ptrpop] {
			return false
		} else {
			stack = stack[:len(stack)-1]
			ptrpop++
		}
	}
	return true
}
