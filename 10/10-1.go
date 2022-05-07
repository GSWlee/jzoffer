package _0

func fib(n int) int {
	nums := []int{0, 1}
	if n < 2 {
		return nums[n]
	}
	for i := 2; i <= n; i++ {
		temp := nums[0]
		nums[0] = nums[1]
		nums[1] = (nums[1] + temp) % (1000000007)
	}
	return nums[1]
}
