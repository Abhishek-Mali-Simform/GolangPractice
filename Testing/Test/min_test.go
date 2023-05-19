package Test

import (
	"Testing/Functions"
	"fmt"
	"testing"
)

func TestMin(test *testing.T) {
	var testCases = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("%d, %d", testCase.a, testCase.b)
		test.Run(testName, func(test *testing.T) {
			ans := Functions.Min(testCase.a, testCase.b)
			if ans != testCase.want {
				test.Errorf("Got %d Want %d", ans, testCase.want)
			}
		})
	}
}
