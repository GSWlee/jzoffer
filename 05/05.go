package main

import "fmt"

func replaceSpace(s string) string {
	replace, length := "%20", len(s)
	new_length := length
	for _, v := range s {
		if v == ' ' {
			new_length = new_length + 2
		}
	}
	target := make([]byte, new_length)
	for i, j := length-1, new_length-1; i >= 0; i-- {
		if s[i] == ' ' {
			for t := len(replace) - 1; t >= 0; t-- {
				target[j] = replace[t]
				j--
			}
		} else {
			target[j] = s[i]
			j--
		}
	}
	return string(target)
}

func main() {
	s := string("We are happy.")
	s = replaceSpace(s)
	fmt.Println(s)
}
