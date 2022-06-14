package main

import (
	"fmt"
	"strings"
)

// Problem Statement
// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

/*
Notes
1- Start with the first array item. Make it the longest common prefix.
2- Compare that prefix with the next array item.
	perform character by character match
	if characters match
		continue to next characters
	else
		update the common prefix till the matched part
	return

What I missed
- Turned out brute force implementation was perfect.

*/

func main() {
	// call function from here
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	if strs[0] == "" {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		if strs[i] == "" {
			return ""
		}

		prefix = matchCharacters(prefix, strs[i])
	}

	return prefix
}

func matchCharacters(prefix, nextString string) string {
	loopLength := len(prefix)

	if len(prefix) > len(nextString) {
		loopLength = len(nextString)
	}

	newPrefix := make([]string, 0)

	for i := 0; i < loopLength; i++ {
		if prefix[i] == nextString[i] {
			newPrefix = append(newPrefix, string(prefix[i]))
			continue
		}

		return strings.Join(newPrefix, "")
	}

	return strings.Join(newPrefix, "")
}

// Challenge:
