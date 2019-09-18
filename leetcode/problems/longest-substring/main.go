package main

import (
	"bytes"
)

func lengthOfLongestSubstring(s string) int {

	return getLength(r, s)

}

func getLength(ls, s string) int {

	r := make(map[string]bool, 0)
	var bf bytes.Buffer

	for i, ch := range s {

		// if not exist
		chStr := string(ch)
		if _, ok := r[chStr]; !ok {
			bf.WriteString(chStr)
			r[chStr] = true
		} else {
			lastStr := s[:i]
			ns := s[i : len(s)-1]
			nr := make(map[string]bool, 0)
			return getLength(nr, ns)
		}

	}

	sub := bf.String()
	return len(sub)
}

func main() {}
