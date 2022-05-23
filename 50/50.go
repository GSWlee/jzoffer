package _0

func firstUniqChar(s string) byte {
	dict := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if v, ok := dict[s[i]]; ok {
			if v != -1 {
				dict[s[i]] = -1
			}
		} else {
			dict[s[i]] = i
		}
	}
	var ans byte = ' '
	num := len(s)
	for k, v := range dict {
		if v != -1 {
			if v < num {
				ans = k
				num = v
			}
		}
	}
	return ans
}
