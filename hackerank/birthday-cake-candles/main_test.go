package main

import (
	"testing"
)

func TestBirthdayCakeCandles(t *testing.T) {

	var test = []struct {
		Heigths []int32
		Output  int32
	}{
		{[]int32{4, 4, 1, 3}, 2},
	}

	for _, tc := range test {
		result := birthdayCakeCandles(tc.Heigths)
		if result != tc.Output {
			t.Errorf("Expected %v, but get :%v", tc.Output, result)
		}
	}

}
