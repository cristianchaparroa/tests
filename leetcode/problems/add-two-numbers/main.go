package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	return fmt.Sprintf("ListNode(Value:%v -> :%v)", n.Val, n.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	firstValue := getTotal(l1)
	secondValue := getTotal(l2)

	fmt.Println(firstValue)
	fmt.Println(secondValue)

	sum := firstValue + secondValue
	fmt.Println(sum)
	rev := revNumber(sum)

	fmt.Println(rev)

	return getNodeReverseNode(rev)
}

func getTotal(l *ListNode) int {
	list := make([]int, 0)
	vals := getValues(l, list)
	return reverse(vals)
}

func GetVals(l *ListNode) []int {
	list := make([]int, 0)
	return getValues(l, list)
}

func getValues(l *ListNode, list []int) []int {

	if l == nil {
		return list
	}

	list = append(list, l.Val)

	return getValues(l.Next, list)
}

func reverList(l []int) []int {
	size := len(l)

	rev := make([]int, 0)

	for i := size - 1; i >= 0; i-- {
		rev = append(rev, l[i])
	}

	return rev
}

func revNumber(number int) []int {
	rev := reverseNumber(number, "")
	return strToListInt(rev)
}

func reverseNumber(number int, partialReverse string) string {
	if number < 10 {
		return fmt.Sprintf("%v%v", partialReverse, number)
	}

	if (number % 10) != 0 {
		mod := number % 10
		pe := number / 10
		e := fmt.Sprintf("%v%v", partialReverse, mod)
		return reverseNumber(pe, e)
	}

	if (number % 10) == 0 {
		pe := number / 10
		e := fmt.Sprintf("%v%v", partialReverse, 0)
		return reverseNumber(pe, e)
	}

	return partialReverse
}

func strToListInt(str string) []int {
	rev := make([]int, 0)

	for _, rune := range str {
		val := strToInt(rune)
		rev = append(rev, val)
	}
	return rev
}

func strToInt(rune rune) int {
	val, err := strconv.Atoi(fmt.Sprintf("%c", rune))

	if err == nil {
		return val
	}
	return -1
}

func reverse(l []int) int {
	fmt.Println(l)
	reverList := reverList(l)

	str := ""
	for _, v := range reverList {
		str = fmt.Sprintf("%v%v", str, v)
	}

	val, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(str)
		return -1
	}

	return val
}

func getNodeReverseNode(rev []int) *ListNode {

	if len(rev) == 0 {
		return nil
	}

	ns := make([]*ListNode, 0)

	nsize := len(rev)
	for i := 0; i < nsize; i++ {
		n := &ListNode{Val: rev[i]}
		ns = append(ns, n)
	}

	var node *ListNode

	for i := 0; i < nsize; i++ {
		cn := ns[i]

		if i == 0 {
			node = cn
		}

		if i < (nsize - 1) {
			nx := i + 1
			nn := ns[nx]
			cn.Next = nn
		}
	}

	return node

}
