package main

func rotateImage(a [][]int) [][]int {
	s := len(a)
	r := zeros(s)
	x := 0
	for j := 0; j < s; j++ {
		y := 0
		for i := (s - 1); i >= 0; i-- {
			r[x][y] = a[i][j]
			y = y + 1
		}
		x = x + 1
	}
	return r
}

func zeros(s int) [][]int {
	z := make([][]int, s)
	for i := 0; i < s; i++ {
		r := make([]int, s)
		z[i] = r
	}
	return z
}
