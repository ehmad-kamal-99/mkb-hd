package main

import (
	"strconv"
)

// Problem Statement
// Given an integer x, return true if x is palindrome integer.

/*
Notes

1- Convert to string
2- Iterate with 2 pointers (start & end) to half point		// didn't work for me
	if pointer.values are equal
		continue
	else
		error

1001 2
1001001 3.5
10101 2.5
111 1.2
11 1
1221 2
13422431 4

What worked was starting pointer at beginning & using that to calculate last index value and comparing.

What I missed
- Understand loops better
- Study loops with two pointers (How to decrement end pointer)

*/

func main() {
	// call function from here
	palindromeNumber(13422431)
}

func palindromeNumber(number int) bool {
	stringForm := strconv.Itoa(number)

	for i := 0; i < len(stringForm)/2; i++ {
		if stringForm[i] != stringForm[len(stringForm)-1-i] {
			return false
		}
	}

	return true
}

// Challenge: Without converting to string.
