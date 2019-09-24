package main

func jumpingOnClouds(c []int32) int32 {

	var jumps int32
	n := len(c)

	for i := 0; i < n; i++ {
		jumps++
		if i < n-2 && c[i+2] == 0 {
			i++
		}
	}
	return jumps - 1
}

func main() {}
