package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/abramtrinh/invoker-combos/pkg/starsBars"
)

const (
	quas  = 1
	wex   = 6
	exort = 7
)

type spell struct {
	spellValue int
	spellName  string
}

// Array/Slice ideal for picking random spell since using maps is not the right DS for the job.
// Not const. Can make it into a func that you return to mimic constant.
var spellArray = [10]spell{
	{spellValue: 3 * quas, spellName: "Cold Snap"},
	{spellValue: 2*quas + wex, spellName: "Ghost Walk"},
	{spellValue: 2*quas + exort, spellName: "Ice Wall"},
	{spellValue: 3 * wex, spellName: "EMP"},
	{spellValue: 2*wex + quas, spellName: "Tornado"},
	{spellValue: 2*wex + exort, spellName: "Alacrity"},
	{spellValue: 3 * exort, spellName: "Sun Strike"},
	{spellValue: 2*exort + quas, spellName: "Forge Spirit"},
	{spellValue: 2*exort + wex, spellName: "Meteor"},
	{spellValue: quas + wex + exort, spellName: "Deafening Blast"},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if err := checkConsts(); err != nil {
		fmt.Println(err)
		return
	}

	var spellMap map[int]string = createSpellMap()

	scanner := bufio.NewScanner(os.Stdin)

	//Closure to keep track of stats
	counter := statTracker()

	//Could make below into a single function.
	for {
		randSpell := getRandomSpell()
		fmt.Printf("Type in %v\n", randSpell.spellName)

		scanner.Scan()
		scanString := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Printf("ERROR: reading standard input: %v", err)
		}

		if scanString == "quit" {
			break
		}

		//If invalid input, just go to the next random spell.
		scanSpellValue, err := parseInput(scanString)
		if err != nil {
			fmt.Println(err)
			continue
		}

		scanSpellName, ok := spellMap[scanSpellValue]
		if !ok {
			fmt.Printf("ERROR: value %v was not found in spellMap\n", scanSpellValue)
			return
		}

		if scanSpellName == randSpell.spellName {
			fmt.Println("Correct.")
			counter(1)
		} else {
			fmt.Println("Incorrect.")
			counter(2)
		}
	}

	fmt.Println("Quit detected: Now quitting.")
	counter(3)
}

// Checks whether or not quas, wex, exort are valid values.
func checkConsts() error {
	if _, err := starsBars.MathCheck(quas, wex, exort); err != nil {
		return fmt.Errorf("Please change quas, wex, and exort consts: %v\n", err)
	}
	return nil
}

// NOTE: Could just use 1 DS.
// Map is ideal for constant lookup instead of iterating like slices/arrays.
func createSpellMap() map[int]string {
	spellMap := make(map[int]string, len(spellArray))
	for _, spell := range spellArray {
		spellMap[spell.spellValue] = spell.spellName
	}
	return spellMap
}

// Retrieves a random spell from spellArray
func getRandomSpell() spell {
	return spellArray[rand.Intn(len(spellArray))]
}

// NOTE: Could be using rune array/slice if I wanted to take care of cases like Turkish font.
// NOTE: To be fair, ranging over string decodes each rune and not byte.
// Parses/validates string of len=3 into the spellValue. ERRORs when non-qwe or len!=3
func parseInput(input string) (int, error) {
	input = strings.ToLower(input)
	var tempValue int = 0
	if len(input) != 3 {
		return tempValue, fmt.Errorf("ERROR: Input of len=%d", len(input))
	}
	for _, value := range input {
		switch value {
		// Raw strings literals
		case 'q':
			tempValue += quas
		case 'w':
			tempValue += wex
		case 'e':
			tempValue += exort
		default:
			return tempValue, fmt.Errorf("ERROR: Found incorrect rune/char of %c", value)
		}
	}
	return tempValue, nil
}

// Closure that keeps track of combo and accuracy stats based on parameter. Can also print current stats.
func statTracker() func(action int) {
	var numCorrect, numTries float32 = 0, 0
	currCombo, maxCombo := 0, 0

	return func(action int) {
		//1 for correct, 2 for incorrect, 3 for stats
		switch action {
		//Right answer, +1 combo, tries, and correct
		case 1:
			currCombo++
			numCorrect++
			numTries++
			//Keep track of highest combo count
			if maxCombo < currCombo {
				maxCombo = currCombo
			}
		//Wrong answer, reset combo counter and +1 tries
		case 2:
			currCombo = 0
			numTries++
		//Prints stats out.
		case 3:
			//Should I check for divide by 0 error?
			fmt.Printf("Combos: %d/%d (Current/Max) | Accuracy: %.4v%%\n", currCombo, maxCombo, numCorrect/numTries*100)
		}
	}

}
