# invoker-combos

> CLI tool for practicing Invoker spell combinations. Initially made to test out math hypothesis.

### Example
```shell
// Tests values of a, b, and, c if they work.
go build .\cmd\sumCombinations\sumCombinations.go
.\sumCombinations.exe (-i file.txt)

// Invoker spell practice tool.
go build .\cmd\invoker\invoker.go 
.\invoker.exe (-m speedTest/timeAttack)
```

## TODO:
* Leaderboards/Progress tracking implementation.
    * Local file / Google Sheets
* Add browser UI instead of just CLI.
    * Make it so that if you press a 4th key, overrides the 1st. Max 3 inputs.
* More modes.
* More unit testing.
* Flag that displays combinations.
* Extend the math portion of combinations with repetition to more than just 3 orbs and 3 slots.
    * NOTE: If I do extend, try using primes numbers as differences.  

## Contents:
1 pkg & 2 cmd

starsBars exports MathCheck for testing my math.  
    Unit tests here.

sumCombinations is the CLI implementation of starsBars.  
    Supports single test or batch file for mass tests using flags.

invoker is a tool to practice invoker spell combos.  
    Multiple modes.

## Why?
This initially started as a shower thought.  
Things like crafting recipes are usually combinations of 2 or more items to create 1.  
One simple way you can check whether a specific recipe is crafted is by using if/switch statements.  
Is there a simpler, more optimized way to check?  
Yes. Assume you have 3 distinct orbs and 3 slots to assign the orbs (repetition allowed). You would have 10 unique spells. This is nCr with repetition. Instead of checking the 27 permutations of the 3 distinct orbs, you can check the 10 unique spells instead. You do this by assigning each orb with a distinct value adding them together to represent the spell pattern.  
	e.g. q=1, w=2, e=5: qwe = 8, qqe = 7, . . .   
The only issue with this is that certain distinct values of qwe (orbs) cause the spells to not have distinct values.  
	e.g. q=113, w=119, e=101: wwe = 339 == qqq  
    I found this issue since I was using the runes of qwe as values.  
After looking at the pattern of errors and testing, I came to a conclusion.  

```
Let a, b, and, c be distinct, positive integers and a < b < c.  
If a, b, and, c passes these conditions:  
2b != a+c   
3b != 2a+c  
3b != 2c+a  
Then each of their combinations (with repetition) will yield distinct values.
```