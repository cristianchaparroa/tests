package main

import (
	"fmt"
	"strconv"
)

const (

	// AM denotes the clock in the morning
	AM string = "AM"
	// PM denotes the clock after the morning
	PM string = "PM"

	MindnightHour12H int = 12
	MoonHour12H      int = 12
)

func isMindnight(hh int, clock string) bool {
	return clock == AM && hh == MindnightHour12H
}

func isBeforeMoon(hhInt int, clock string) bool {
	return hhInt < (MoonHour12H) && clock == AM || hhInt == MoonHour12H && clock == PM
}

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	clock := s[len(s)-2:]

	hh := s[0:2]
	hhInt, _ := strconv.Atoi(hh)

	mmSSStr := s[2 : len(s)-2]
	if isMindnight(hhInt, clock) {
		return fmt.Sprintf("00%s", mmSSStr)
	}

	if isBeforeMoon(hhInt, clock) {
		return s[:len(s)-2]
	}

	hhStr := hhInt + 12

	return fmt.Sprintf("%v%v", hhStr, mmSSStr)
}

func main() {

}
