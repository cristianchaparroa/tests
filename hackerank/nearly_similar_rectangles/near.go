package near


// a/c = b/d
// d/c = b/a
func NearlySimilarRectangles(sides [][]int64) int64 {
	// Write your code here

	matches := int64(0)

	for i := 0; i < len(sides)-1; i++ {
		side := sides[i]
		a := side[0]
		b := side[1]

		for j := i + 1; j < len(sides); j++ {

			nextSide := sides[j]
			c := nextSide[0]
			d := nextSide[1]

			if a*d == b*c {
				matches++
			}
		}
	}
	return matches
}
