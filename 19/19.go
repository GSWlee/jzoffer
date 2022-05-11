package main

//
//  isMatch
//  @Description: 递归
//  @param s
//  @param p
//  @return bool
//
func isMatchv1(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) != 0 && len(p) == 0 {
		return false
	}
	if len(p) < 2 {
		if len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
			return isMatch(s[1:], p[1:])
		} else {
			return false
		}
	} else {
		if p[1] == '*' {
			if len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
				return isMatch(s[1:], p[2:]) || isMatch(s[1:], p) || isMatch(s, p[2:])
			} else {
				return isMatch(s[:], p[2:])
			}
		} else {
			if len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
				return isMatch(s[1:], p[1:])
			} else {
				return false
			}
		}
	}
}

//
//  isMatch
//  @Description: dp
//  @param s
//  @param p
//  @return bool
//
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

func main() {
	isMatch("ab", ".*c")
}
