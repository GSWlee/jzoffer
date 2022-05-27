package _8

func reverseWords(s string) string {
	ans := ""
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if start != i {
				ans = s[start:i] + " " + ans
			}
			start = i + 1
		}
	}
	if start < len(s) {
		ans = s[start:] + " " + ans
	}
	if len(ans) > 0 {
		return ans[:len(ans)-1]
	}
	return ans
}
