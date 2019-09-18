package main

func groupSocks(ar []int32) map[int32]int32 {
	pairs := make(map[int32]int32, 0)

	for _, v := range ar {
		val, found := pairs[v]

		if found {
			pairs[v] = val + 1
		} else {
			pairs[v] = 1
		}

	}
	return pairs
}

func isEven(n int32) bool {
	return n%2 == 0
}

func countPairs(p map[int32]int32) int32 {

	var pairs int32

	for _, v := range p {

		if isEven(v) {
			pairs = pairs + v/2
		} else {
			v = (v - 1)
			pairs = pairs + v/2
		}
	}

	return pairs
}

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	pairs := groupSocks(ar)
	return countPairs(pairs)
}

func main() {}
