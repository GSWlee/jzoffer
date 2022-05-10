package _6

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if x == 0.0 {
		return 0.0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	ans := 1.0
	temp := myPow(x, n/2)
	if n%2 == 1 {
		ans = temp * temp * x
	} else {
		ans = temp * temp
	}
	return ans
}
