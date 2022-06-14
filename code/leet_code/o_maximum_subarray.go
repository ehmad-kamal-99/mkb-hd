package main

import (
	"fmt"
)

// Problem Statement
// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// A subarray is a contiguous part of an array.

/*
Notes

Brute Force way with O(n2)


*/

func main() {
	// call function from here
	fmt.Println(maximumSubarray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maximumSubarray([]int{1}))
	fmt.Println(maximumSubarray([]int{5, 4, -1, 7, 8}))
}

func maximumSubarray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	finalSum := 0

	for i := 0; i < len(nums); i++ {
		localSum := nums[i]

		for j := i + 1; j < len(nums); j++ {
			localSum += nums[j]
			if localSum > finalSum {
				finalSum = localSum
			}
		}
	}

	return finalSum
}

func maximumSubarrayON(nums []int) int {
	return 0
}

// Challenge: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
