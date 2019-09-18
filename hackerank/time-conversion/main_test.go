package main

import "testing"

func TestTimeConversion(t *testing.T) {

	var test = []struct {
		Input          string
		ExpectedOutput string
	}{
		{"07:05:45PM", "19:05:45"},
		{"12:05:45PM", "12:05:45"},
		{"12:03:45AM", "00:03:45"},
		{"03:05:45PM", "15:05:45"},
	}

	for _, tc := range test {
		result := timeConversion(tc.Input)

		if tc.ExpectedOutput != result {
			t.Errorf("Expect %s, but get: %s", tc.ExpectedOutput, result)
		}
	}
}
