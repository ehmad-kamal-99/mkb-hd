package main

import (
	"fmt"
)

// Problem Statement
// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.

/*
Notes



*/

func main() {
	// call function from here
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	verifyMap := make(map[int]int)
	verifyMap[0] = 0
	verifyMap[1] = 0
	verifyMap[2] = 0
	verifyMap[3] = 0
	verifyMap[4] = 0
	verifyMap[5] = 0
	verifyMap[6] = 0
	verifyMap[7] = 0
	verifyMap[8] = 0
	verifyMap[9] = 0

	for _, num := range nums {
		verifyMap[num]++
	}

	for key, _ := range verifyMap {
		if verifyMap[key] == 1 {
			return key
		}
	}

	return 0
}

// Challenge:
