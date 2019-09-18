package main

import (
	"testing"
)

func TestGetSum(t *testing.T) {

	a := &ListNode{Val: 2}
	b := &ListNode{Val: 4}
	c := &ListNode{Val: 3}

	a.Next = b
	b.Next = c

	initValue := make([]int, 0)
	values := getValues(a, initValue)
	expectedSize := 3
	resSize := len(values)

	if expectedSize != resSize {
		t.Errorf("Expected the following size: %v, but get:%v", expectedSize, resSize)
	}
}

func TestReverList(t *testing.T) {

	toRev := []int{1, 2, 3}
	rev := reverList(toRev)
	AssertTrue(t, 3, rev[0])
	AssertTrue(t, 2, rev[1])
	AssertTrue(t, 1, rev[2])
}

func AssertTrue(t *testing.T, expected, value int) {
	if expected != value {
		t.Errorf("Expected :%v, but get:%v", expected, value)
	}
}

func TestReverseNumber(t *testing.T) {

	var test = []struct {
		Number   int
		Expected string
	}{
		{807, "708"},
	}

	for _, ts := range test {
		res := reverseNumber(ts.Number, "")

		if ts.Expected != res {
			t.Errorf("Expected %v, but get:%v", ts.Expected, res)
		}
	}
}

func TestReverse(t *testing.T) {
	var test = []struct {
		List     []int
		Expected int
	}{
		{[]int{1, 2, 3}, 321},
		{[]int{4, 7, 8}, 874},
	}

	for _, ts := range test {
		rev := reverse(ts.List)

		AssertTrue(t, ts.Expected, rev)
	}
}

func TestGetNodeReverseNode(t *testing.T) {
	//list := []int{3, 2, 1}
	// node := getNodeReverseNode(list)
}

func TestAddTwoNumbers(t *testing.T) {
	a := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	b := []int{5, 6, 4}
	result := []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	aList := getNodeReverseNode(a)
	bList := getNodeReverseNode(b)

	expected := getNodeReverseNode(result)

	resList := addTwoNumbers(aList, bList)

	if !isListEqual(expected, resList) {
		t.Errorf("\nExpected:\n\t%v, \nbut get:\n\t%v", expected, resList)
	}
}

func isListEqual(a, b *ListNode) bool {
	aList := GetVals(a)
	bList := GetVals(b)

	if len(aList) != len(bList) {
		return false
	}

	for i := 0; i < len(aList); i++ {
		av := aList[i]
		bv := bList[i]

		if av != bv {
			return false
		}
	}

	return true
}
