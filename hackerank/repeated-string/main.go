package main

import (
	"math"
)

const (
	ToSearch string = "a"
)

func countOcurencies(s string) int64 {

	var count int64

	for _, c := range s {
		if ToSearch == string(c) {
			count++
		}
	}

	return count

}

func getSubString(s string, n int64) string {
	l := int64(len(s))
	rest := n % l

	isComplete := rest == 0
	subString := ""
	if !isComplete {
		subString = s[:rest]
	}

	return subString
}

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {

	o := countOcurencies(s)
	ss := getSubString(s, n)
	oss := countOcurencies(ss)
	rep := n / int64(len(s))
	fr := math.Floor(float64(rep))

	return (o * int64(fr)) + oss
}

func main() {}
