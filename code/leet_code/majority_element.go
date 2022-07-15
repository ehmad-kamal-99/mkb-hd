package main

import (
	"fmt"
)

// Problem Statement
// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

/*
Notes



*/

func main() {
	// call function from here
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	count, candidate := 0, 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

// Challenge: Could you solve the problem in linear time and in O(1) space?
