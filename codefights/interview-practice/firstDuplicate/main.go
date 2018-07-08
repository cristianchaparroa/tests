package main

import (
	"fmt"
	"time"
)

/*
func firstDuplicate(a []int) int {

    limit := 100000
    idxs  := limit
    size  := len(a)

    if ( size == 0 ) || ( size == 1 ) || (size > limit)  {
        return -1
    }

    for i, cn := range a {
        for j := i+1; j < size; j++ {
            if ((i != j) && (cn == a[j]) && (j < idxs)) {
                idxs = j
                continue;
            }
        }
    }

     O(n^2)
    if (idxs <= size) {
        return a[idxs]
    }

    return -1
}
*/

func firstDuplicate(a []int) int {
	size := len(a)
	m := make([]bool, size)

	for _, _ = range a {
		m = append(m, false)
	}
	// O(N)
	for _, v := range a {
		if m[v] == false {
			m[v] = true
		} else {
			return v
		}
	}
	// O(N)
	// O(2N) -> O(N)
	return -1
}
