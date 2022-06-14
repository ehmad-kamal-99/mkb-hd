package main

import (
	"fmt"
)

// Problem Statement
// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The relative order of the elements may be changed.
// Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
// Return k after placing the final result in the first k slots of nums.
// Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

/*
Notes

Seems simple. Start iterating over array.
If value == required val, update it to 51
move to next value until all array traversed
sort it, and return count of value that aren't 51.

*/

func main() {
	// call function from here
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	if len(nums) < 2 {
		if len(nums) == 0 {
			return 0
		}

		if nums[0] == val {
			return 0
		}
	}

	convertCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = 51
			convertCount++
		}
	}

	sortArray2(nums)

	return len(nums) - convertCount
}

func sortArray2(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
			}
		}
	}

	return nums
}

// Challenge:
