package starsBars

import (
	"errors"
	"fmt"
)

//custom type errors with Error() method or errors.new("") or fmt.Errorf()
//thoughts on returning just (error) instead of something like (int[], error)
//best practices
//when do people use fatal or logs hmm
//hmm best practices for writing unit tests like table driven.
//i could use abstractions here or interface but i will 100% forget
// reason for returning bool, error is for table driven uunit testing (ami doing it right)

// Checks for a, b, c to be distinct values.
func distinct(a, b, c int) (bool, error) {
	if a == b || a == c || b == c {
		//return errors.New("values are not distinct")
		return false, fmt.Errorf("%d, %d, %d are not distinct", a, b, c)
	}
	return true, nil
}

// Checks for a, b, c to be non-zero positive numbers.
func positive(a, b, c int) (bool, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		//return errors.New("at least one value is not positive")
		return false, fmt.Errorf("%d, %d, %d are not non-zero positive", a, b, c)
	}
	return true, nil
}

// Sorts a, b, c such that a < b < c.
func sortAscend(a, b, c int) (newA, newB, newC int) {
	if a > c {
		c, a = a, c
	}
	if a > b {
		b, a = a, b
	}
	if b > c {
		c, b = b, c
	}
	return a, b, c
}

// Checks for a, b, c such that no combinations have the same sum.
// This function cuts down the need to check all 10 sum combinations of abc.
// Refer to writeup attached for explanation.
// TODO: WriteUp
func repeatedSum(a, b, c int) (bool, error) {
	// switch is used so I can practice using the switch{}
	// can be rewritten as if statements
	// NOTE: when to use switch over if statements (effective go)
	// maybe use error wrapping here. fmt.Errorf()
	switch {
	case 2*b == (a + c):
		return false, errors.New("the 1 1 test failed")
	case 3*b == (2*a + c):
		return false, errors.New("the 1 2 test failed")
	case 3*b == (2*c + a):
		return false, errors.New("the 2 1 test failed")
	}
	return true, nil
}

// Exported
// Checks for a, b, c such that they are distinct, non-zero, positive, sorted, and have no repeated sums.
func MathCheck(a, b, c int) (bool, error) {
	if succeed, err := distinct(a, b, c); err != nil {
		return succeed, err
	}

	if succeed, err := positive(a, b, c); err != nil {
		return succeed, err
	}

	sortedA, sortedB, sortedC := sortAscend(a, b, c)

	if succeed, err := repeatedSum(sortedA, sortedB, sortedC); err != nil {
		return succeed, err
	}

	return true, nil
}
