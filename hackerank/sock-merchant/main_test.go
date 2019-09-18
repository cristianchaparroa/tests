package main

import "testing"

func TestSockMerchant(t *testing.T) {

	var test = []struct {
		n              int32
		ar             []int32
		expectedResult int32
	}{
		{9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}, 3},
		{2, []int32{1, 1}, 1},
		{3, []int32{1, 1, 3}, 1},
		{3, []int32{1, 1, 1}, 1},
	}

	for _, tc := range test {
		result := sockMerchant(tc.n, tc.ar)

		if tc.expectedResult != result {
			t.Errorf("Expected %v, but get:%v", tc.expectedResult, result)
		}
	}

}
