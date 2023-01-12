package starsBars

import (
	"fmt"
	"testing"
)

type testStruct struct {
	a, b, c      int
	expectedBool bool
}

func TestDistinct(t *testing.T) {
	// So if I have a func that only returns error, using table-driven testing would be a pain.
	// e.g. I want 123 to return nil (no error) and 111 to return an error
	// So in a struct, I can set errorExist := nil and that would work for checking 123
	// But how do I make it so that the errorExists == when 111 returns an error
	// Solution 1 is using regex to string match I guess? Solution 2 is just use another value that can soley be used for testing.

	testTable := []testStruct{
		{1, 1, 1, false},
		{1, 1, 2, false},
		{1, 2, 3, true},
	}

	// NOTE: Used t.Run() because it lets me "see" "subtests" from the table-driven testing.
	// NOTE: Tested with outputBool and expectedBool instead of err !=/== nil because I want to see some fail on purpose.
	for _, test := range testTable {
		testName := fmt.Sprintf("%d %d %d", test.a, test.b, test.c)
		t.Run(testName, func(t *testing.T) {
			outputBool, _ := distinct(test.a, test.b, test.c)
			if outputBool != test.expectedBool {
				t.Errorf("%d %d %d returned %t, expected %t ", test.a, test.b, test.c, outputBool, test.expectedBool)
			}
		})
	}
}

func TestPositive(t *testing.T) {
	testTable := []testStruct{
		{1, 1, 1, true},
		{1, 1, -1, false},
		{1, 0, -1, false},
	}

	for _, test := range testTable {
		testName := fmt.Sprintf("%d %d %d", test.a, test.b, test.c)
		t.Run(testName, func(t *testing.T) {
			outputBool, _ := positive(test.a, test.b, test.c)
			if outputBool != test.expectedBool {
				t.Errorf("%d %d %d returned %t, expected %t ", test.a, test.b, test.c, outputBool, test.expectedBool)
			}
		})
	}
}

func TestSortAscend(t *testing.T) {
	var testTable = []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{1, 3, 2},
		{3, 2, 1},
		{3, 1, 2},
		{2, 1, 3},
		{2, 3, 1},
	}

	sortedA, sortedB, sortedC := 1, 2, 3

	for _, test := range testTable {
		testName := fmt.Sprintf("%d%d%d", test.a, test.b, test.c)
		t.Run(testName, func(t *testing.T) {
			outputA, outputB, outputC := sortAscend(test.a, test.b, test.c)
			if (sortedA != outputA) || (sortedB != outputB) || (sortedC != outputC) {
				t.Errorf("%d%d%d did not get sorted, returned %d%d%d", test.a, test.b, test.c, outputA, outputB, outputC)
			}
		})
	}
}

func TestRepeatedSum(t *testing.T) {
	testTable := []testStruct{
		{1, 2, 3, false},
		{1, 3, 4, false},
		{1, 2, 4, false},
		{1, 6, 13, true},
	}

	for _, test := range testTable {
		testName := fmt.Sprintf("%d %d %d", test.a, test.b, test.c)
		t.Run(testName, func(t *testing.T) {
			outputBool, _ := repeatedSum(test.a, test.b, test.c)
			if outputBool != test.expectedBool {
				t.Errorf("%d %d %d returned %t, expected %t ", test.a, test.b, test.c, outputBool, test.expectedBool)
			}
		})
	}
}

func TestMathCheck(t *testing.T) {
	testTable := []testStruct{
		{1, 1, 3, false},
		{1, -3, 2, false},
		{1, 3, 5, false},
		{1, 8, 13, true},
	}

	// NOTE: I could make this below into a function since I reimplemented this 4x already. But might hurt readbility and future modifications.
	for _, test := range testTable {
		testName := fmt.Sprintf("%d %d %d", test.a, test.b, test.c)
		t.Run(testName, func(t *testing.T) {
			outputBool, _ := MathCheck(test.a, test.b, test.c)
			if outputBool != test.expectedBool {
				t.Errorf("%d %d %d returned %t, expected %t ", test.a, test.b, test.c, outputBool, test.expectedBool)
			}
		})
	}
}
