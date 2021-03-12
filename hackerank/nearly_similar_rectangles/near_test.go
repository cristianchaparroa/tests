package near

import (
	"fmt"
	"testing"
)

func TestNearlySimilarRectangles(t *testing.T) {
	var test = []struct {
		input          [][]int64
		expectedResult int64
	}{
		{[][]int64{{2, 1}, {10, 7}, {9, 5}, {6, 9}, {7, 3}}, int64(0)},
		{[][]int64{{5, 10}, {10, 10}, {3, 6}, {9, 9}}, int64(2)},
	}

	for _, t := range test {
		result := NearlySimilarRectangles(t.input)

		if t.expectedResult != result {
			message := fmt.Errorf("Expected %v, but get:%v", t.expectedResult, result)
			fmt.Println(message)
		}
	}
}
