package Test

import (
	"Testing/Functions"
	"fmt"
	"testing"
)

type TestCase[Num int | float64] struct {
	typeUsed string
	num      Num
	want     Num
}

func getTestCase[Num int | float64](test *testing.T, testCase TestCase[Num]) {
	test.Helper()
	testName := fmt.Sprintf("%s, %v", testCase.typeUsed, testCase.num)
	test.Run(testName, func(test *testing.T) {
		ans := Functions.Factorial(testCase.num)
		if ans != testCase.want {
			test.Errorf("got %v want %v", ans, testCase.want)
		}
	})
}

func TestFactorial(test *testing.T) {
	getTestCase(test, TestCase[int]{
		typeUsed: "int",
		num:      5,
		want:     120,
	})
	getTestCase(test, TestCase[float64]{
		num:      5.5,
		typeUsed: "float64",
		want:     287.885277815,
	})
}
