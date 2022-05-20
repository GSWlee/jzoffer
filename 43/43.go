package _3

func countDigitOne(n int) int {
	if n == 0 {
		return 0
	}
	tag, w := 1, 1
	for 10*tag <= n {
		tag = tag * 10
		w++
	}
	if w == 1 {
		return 1
	}
	ys := n % tag
	s := (n - ys) / tag
	num_s := s * (w - 1) * tag / 10
	num_y := 0
	if s > 1 {
		num_y = tag
	} else {
		num_y = ys + 1
	}
	return countDigitOne(ys) + num_s + num_y
}
