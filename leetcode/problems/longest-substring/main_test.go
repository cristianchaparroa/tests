package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	var test = []struct {
		Test           string
		ExpectedLength int
	}{
		/*	{"abcabcbb", 3},*/
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for i, tc := range test {
		resultLength := lengthOfLongestSubstring(tc.Test)

		if resultLength != tc.ExpectedLength {
			t.Errorf("Case-%v:Expected %v, but get:%v\n", (i + 1), tc.ExpectedLength, resultLength)
		}
	}
}
