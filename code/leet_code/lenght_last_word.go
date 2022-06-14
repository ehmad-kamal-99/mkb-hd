package main

import (
	"fmt"
	"strings"
)

// Problem Statement
// Given a string s consisting of words and spaces, return the length of the last word in the string.
// A word is a maximal substring consisting of non-space characters only.

/*
Notes

First, start with trimming spaces.
Then, split based on spaces and pick the last one and return its length.

*/

func main() {
	// call function from here
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)

	stringArr := strings.Split(s, " ")

	return len(stringArr[len(stringArr)-1])
}

// Challenge:
