package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/abramtrinh/invoker-combos/pkg/starsBars"
)

var inputFlag = flag.String("i", "", "input text file name as string")

// os.Exit(1) vs return vs continue vs break
// NOTE: Using return instead of os.Exit(1) is more "graceful?"
func main() {

	flag.Parse()

	// -i flag used
	if *inputFlag != "" {
		//Opens file
		file, err := os.Open(*inputFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		//Parses file line by line and appends to a string slice
		fileScanner := bufio.NewScanner(file)
		var fileLines []string
		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}

		if err := fileScanner.Err(); err != nil {
			fmt.Printf("file scanner errored: %v\n", err)
			os.Exit(1)
		}

		//Iterates through string slice, converts to int array, and MathCheck those values.
		for _, line := range fileLines {
			abcValue := strings.Fields(line)
			intArray, err := strSlice2IntArray(abcValue)
			if err != nil {
				fmt.Printf("%v failed: %v\n", abcValue, err)
				continue
			}
			_, err = starsBars.MathCheck(intArray[0], intArray[1], intArray[2])
			if err != nil {
				fmt.Printf("%v failed: %v\n", intArray, err)
				continue
			}
			fmt.Printf("%d %d %d passes the starBars tests.\n", intArray[0], intArray[1], intArray[2])
		}
		return
	}

	//Instead of flag pkg for this, could use os.Args[1:] (would have to modify switch)
	switch len(flag.Args()) {
	case 0:
		fmt.Println("Please input 3 positive and distinct numbers seperated by spaces.")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Printf("ERROR: reading standard input: %v", err)
		}
		words := strings.Fields(text)
		intArray, err := strSlice2IntArray(words)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = starsBars.MathCheck(intArray[0], intArray[1], intArray[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%d %d %d passes the starBars tests.\n", intArray[0], intArray[1], intArray[2])
	case 3:
		intArray, err := strSlice2IntArray(flag.Args())
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = starsBars.MathCheck(intArray[0], intArray[1], intArray[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%d %d %d passes the starBars tests.\n", intArray[0], intArray[1], intArray[2])
	default:
		fmt.Println("Please rerun this with either 3 numbers or none.")
		os.Exit(1)
	}
	os.Exit(0)
}

// Converts string slice to an int array of size 3.
func strSlice2IntArray(strSlice []string) ([3]int, error) {
	//Could use slice instead of array too here.
	intArray := [3]int{-1, -1, -1}
	if len(strSlice) != 3 {
		return intArray, fmt.Errorf("%d inputs found, need 3", len(strSlice))
	}
	for i, strValue := range strSlice {
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			//Reach here when you try to convert non-int string
			return intArray, fmt.Errorf("found non-int %s", strValue)
		}
		intArray[i] = intValue
	}
	return intArray, nil
}
