package _5

func add(a int, b int) int {
	sum := a ^ b
	carry := a & b << 1
	for carry != 0 {
		a, b = sum, carry
		sum = a ^ b
		carry = a & b << 1
	}
	return sum
}
