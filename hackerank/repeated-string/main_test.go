package main

import "testing"

func TestRepeatedString(t *testing.T) {

	var test = []struct {
		n           int64
		s           string
		repExpected int64
	}{
		{int64(731869010806), "aedbdaeaddadddcedcbbabdccbecaecaccdbebeeadadcaabbaabbaeeeecaddbcdecbbdccdebaaebecdaaabbcdeccbabaabce", 168329872486},
	}

	for _, tc := range test {
		result := repeatedString(tc.s, tc.n)

		if tc.repExpected != result {
			t.Errorf("Expected %v, but get:%v", tc.repExpected, result)
		}
	}
}
