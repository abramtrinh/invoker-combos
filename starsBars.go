package main

import (
	"errors"
	"fmt"
)

//custom type errors with Error() method or errors.new("")

// TODO: UnitTesting 111 123 112
// Checks for a, b, c to be distinct values.
func distinct(a, b, c int) error {
	if a == b || a == c || b == c {
		return errors.New("values are not distinct")
	}
	return nil
}

// TODO: UnitTesting -1, 0, 1
// Checks for a, b, c to be non-zero positive numbers.
func positive(a, b, c int) error {
	if a <= 0 || b <= 0 || c <= 0 {
		return errors.New("at least one value is not positive")
	}
	return nil
}

// TODO: UnitTesting 123 132 321 312 213 231
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

// TODO: UnitTesting 11 12 21 (fails) and a succeeding case
// Checks for a, b, c such that no combinations have the same sum.
// This function cuts down the need to check all 10 sum combinations of abc.
// Refer to writeup attached for explanation.
// TODO: WriteUp
func repeatedSum(a, b, c int) error {
	//switch is used so I can practice using the switch{}
	//can be rewritten as if statements
	//note: when to use switch over if statements (effective go)
	switch {
	case 2*b == (a + c):
		return errors.New("fails the 1 1 test")
	case 3*b == (2*a + c):
		return errors.New("fails the 1 2 test")
	case 3*b == (2*c + a):
		return errors.New("fails the 2 1 test")
	}
	return nil
}

// TODO: UnitTesting
func mathCheck(a, b, c int) error {
	if err := distinct(a, b, c); err != nil {
		return err
	}

	if err := positive(a, b, c); err != nil {
		return err
	}

	sortedA, sortedB, sortedC := sortAscend(a, b, c)

	if err := repeatedSum(sortedA, sortedB, sortedC); err != nil {
		return err
	}

	return nil
}

// split this project. 1. starBars 2.unit testing added 3.flags functionality 4. modality 5. cleanjuop
func main() {
	// File is named starsBars because it is a reference to the stars and bars method in combinatorics.

	// TODO: Test combinations of 11 12 21 and something that is that not those.

	// TODO: a, b, c values where they are invalid (you know they are invalid since they are still the default value of -1)
	// To explain clearly, entries like aa ss dd or like aa 2 3 should fail.
	// So throw error when a, b, c == -1 here. Or just move next.

	/*
		//need to check for valid filename like foo.txt and othing else. errorcheck i think
		//if -i -o are nil that means do not use those flags i guess. then error check for bad filenames
		fInputPtr := flag.String("i", "nil", " string of input filename")
		fOutputPtr := flag.String("o", "nil", "string of output filename")
		flag.Parse()

		fmt.Println("-i", *fInputPtr)
		fmt.Println("-o", *fOutputPtr)
		//can use flag.Args() for the trailing 3 numbers and do a single run like snb.go 1 33 22, if not size 3 .
		//highkey might just ignore this until last. just focus on normal prompts
		fmt.Println("tail:", flag.Args())
		//if no i o or args then just run normal prompt
		//just use if statesments for each flag . liek setup with if statements
		//4 styles. -i -o flag.Args() or just normal prompt. thats as an exe. now when i export as package i need a function that does the same but with 3 params.
	*/

	/*
		fmt.Scanf("%d %d %d\n", &a, &b, &c)
		fmt.Printf("recieved\n")
	*/

	//io use defer to close.

	var a, b, c int = 1, 3, 8
	fmt.Printf("%d %d %d\n", a, b, c)
	if err := mathCheck(a, b, c); err != nil {
		fmt.Print(err)
	}
	fmt.Print("it works")

}
