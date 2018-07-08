package main

import (
	"reflect"
	"testing"
)

func TestRotateImage(t *testing.T) {
	var test = []struct {
		input    [][]int
		expected [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{[][]int{{1}}, [][]int{{1}}},
		{[][]int{{10, 9, 6, 3, 7}, {6, 10, 2, 9, 7}, {7, 6, 3, 8, 2}, {8, 9, 7, 9, 9}, {6, 8, 6, 8, 2}}, [][]int{{6, 8, 7, 6, 10}, {8, 9, 6, 10, 9}, {6, 7, 3, 2, 6}, {8, 9, 8, 9, 3}, {2, 9, 2, 7, 7}}},
	}
	for _, ts := range test {
		r := rotateImage(ts.input)
		if !areArraysEquals(r, ts.expected) {
			t.Errorf("expected:%v, but get:%v", ts.expected, r)
		}
	}
}
// It verifies if a and b arrays contains the same values
// Returns the validation status.
func areArraysEquals(a, b [][]int) bool {
	return reflect.DeepEqual(a, b)
}
