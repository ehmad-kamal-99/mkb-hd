package main

import (
	"fmt"
	"strings"
)

// Problem Statement
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.

/*
Notes



*/

func main() {
	// call function from here
	fmt.Println(validPalindrome("A man, a plan, a canal: Panama"))
}

func validPalindrome(s string) bool {
	strippedString := strings.ToLower(strip([]byte(s)))

	if len(strippedString) < 2 {
		return true
	}

	return isPalindrome(strippedString)
}

func strip(bytes []byte) string {
	n := 0
	for _, byte := range bytes {
		if ('a' <= byte && byte <= 'z') ||
			('A' <= byte && byte <= 'Z') ||
			('0' <= byte && byte <= '9') {
			bytes[n] = byte
			n++
		}
	}
	return string(bytes[:n])
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for i := 0; i < len(s)/2; i++ {
		if s[start] == s[end] {
			start++
			end--

			continue
		}

		return false
	}

	return true
}

// Challenge:
