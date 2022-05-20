package _4

func findNthDigit(n int) int {
	now, tag, w := 10, 10, 1
	n = n + 1
	for now <= n {
		w++
		n = n - now
		now = w * 9 * tag
		tag = tag * 10
	}
	tag = tag / 10
	if tag == 1 {
		tag--
	}
	ys := n % w
	s := (n - ys) / w
	if ys == 0 {
		return (tag + s - 1) % 10
	}

	tag = tag + s
	ws := w - ys + 1

	for i := 1; i < ws; i++ {
		tag = (tag - (tag % 10)) / 10
	}
	return tag % 10
}
