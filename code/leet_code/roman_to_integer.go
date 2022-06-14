package main

import "fmt"

// Problem Statement
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//
//For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
//
//Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
//    I can be placed before V (5) and X (10) to make 4 and 9.
//    X can be placed before L (50) and C (100) to make 40 and 90.
//    C can be placed before D (500) and M (1000) to make 400 and 900.
//
//Given a roman numeral, convert it to an integer.

/*
Notes

Create a dictionary having values against roman numerals
Iterate over each numeral
If the numeral is I, X, C, check that the next numeral isn't V, X or L, C or D, M.
	if true
	combine the 2 and increment the counter by 1
	if false
	use its value from dictionary

What I missed
- Keep track of continue statements
- When copy pasting, make sure to keep a close eye on edge cases
- If a program ends up being big, try running basic program with minimal test case first
*/

func main() {
	// call function from here
	fmt.Println(romanToInteger("III"))
}

func romanToInteger(s string) int {
	romanDictionary := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	numeralValue := make([]int, 0)

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			numeralValue = append(numeralValue, romanDictionary[string(s[i])])
			continue
		}

		if string(s[i]) == "I" || string(s[i]) == "X" || string(s[i]) == "C" {
			if string(s[i]) == "I" {
				if string(s[i+1]) == "V" {
					numeralValue = append(numeralValue, 4)
					i++
					continue
				} else if string(s[i+1]) == "X" {
					numeralValue = append(numeralValue, 9)
					i++
					continue
				}

				numeralValue = append(numeralValue, romanDictionary[string(s[i])])
				continue
			} else if string(s[i]) == "X" {
				if string(s[i+1]) == "L" {
					numeralValue = append(numeralValue, 40)
					i++
					continue
				} else if string(s[i+1]) == "C" {
					numeralValue = append(numeralValue, 90)
					i++
					continue
				}

				numeralValue = append(numeralValue, romanDictionary[string(s[i])])
				continue
			} else if string(s[i]) == "C" {
				if string(s[i+1]) == "D" {
					numeralValue = append(numeralValue, 400)
					i++
					continue
				} else if string(s[i+1]) == "M" {
					numeralValue = append(numeralValue, 900)
					i++
					continue
				}

				numeralValue = append(numeralValue, romanDictionary[string(s[i])])
				continue
			}
		}

		numeralValue = append(numeralValue, romanDictionary[string(s[i])])
	}

	finalInt := 0

	for _, val := range numeralValue {
		finalInt += val
	}

	return finalInt
}

// Challenge:
