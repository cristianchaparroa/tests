package main

import (
	"fmt"
	"sort"
)

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {

	heigths := make(map[int32]int32, 0)

	for _, h := range ar {
		v, found := heigths[h]

		if found {
			heigths[h] = v + 1
		} else {
			heigths[h] = 1
		}
	}

	values := make([]int32, len(heigths))

	for _, v := range heigths {
		values = append(values, v)
	}

	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })

	fmt.Println(values)

	if len(values) == 0 {
		return 0
	}

	return values[0]
}

func main() {}
