package main

import (
	"fmt"
)

// Problem Statement
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

/*
Notes



*/

func main() {
	// call function from here
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("car", "rat"))
}

func isAnagram(s string, t string) bool {
	check := make(map[rune]int)

	for _, i := range s {
		check[i]++
	}

	for _, j := range t {
		if _, ok := check[j]; !ok {
			return false
		}

		check[j]--
	}

	for _, value := range check {
		if value < 0 {
			return false
		}
	}

	return true
}

// Challenge:
