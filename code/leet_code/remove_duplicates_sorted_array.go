package main

import (
	"fmt"
)

// Problem Statement
// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.
// Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
// Return k after placing the final result in the first k slots of nums.
// Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

/*
Notes

Check if length == 1, return 1.
If greater than one, start looping over array,
checkNext 1st & 2nd item
	if equal
		check next, if equal, keep checking until not equal
		if not equal, clean up


*/

// []int{1, 1, 1, 1, 1, 1, 5, 5, 5, 6, 7, 7, 8, 9, 10}

func main() {
	// call function from here
	fmt.Println(removeDuplicatesSortedArray([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 5, 5, 5, 5, 5}))
}

// func removeDuplicatesSortedArray(nums []int) int {
// 	if len(nums) < 2 {
// 		return 1
// 	}
//
// 	fmt.Println("length of og array: ", len(nums))
//
// 	newIndex := 0
// 	var check = false
//
// 	for i := 0; i < len(nums); {
// 		newIndex, check = checkNext(nums, i, i+1)
//
// 		fmt.Println("newIndex: ", newIndex, ", check: ", check)
//
// 		if check {
// 			break
// 		}
// 		if newIndex+1 == len(nums) {
// 			break
// 		}
//
// 		i = newIndex
// 	}
//
// 	fmt.Println(nums)
//
// 	return 0
// }
//
// func checkNext(nums []int, i, j int) (int, bool) {
// 	fmt.Println("initial index: ", i)
// 	fmt.Println("second index: ", j)
//
// 	fmt.Println("comparing: ", nums[i], nums[j])
// 	if nums[i] == nums[j] {
// 		fmt.Println("equal")
// 		nums[j] = -101 // update the duplicate
//
// 		fmt.Println("second index updated to: ", j+1)
// 		j++ // increment j so next item in array gets compared
//
// 		if len(nums) == j { // if j == length of array, array is all traversed.
// 			fmt.Println("length: ", len(nums), " second index: ", j)
// 			return j - 1, true // return j-1 which is the last index and return true to denote that all comparisons are complete.
// 		}
//
// 		checkNext(nums, i, j)
// 	} else {
// 		fmt.Println("not equal, returning new index: ", j)
// 		return j, false
// 	}
//
// 	return j, true
// }

func removeDuplicatesSortedArray(nums []int) int {
	if len(nums) < 2 {
		return 1
	}

	var check = true

	for i := 0; i < len(nums)-1; {
		j := i + 1
		check = true

		for check {
			if j, check = checkNext(nums, i, j); check {
				j++

				if j == len(nums) {
					break
				}
			} else {
				check = false
				i = j
			}
		}

		if j == len(nums) {
			break
		}
	}

	nums = sortArray(nums)

	counter := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 101 {
			break
		}

		counter++
	}

	return counter
}

func checkNext(nums []int, initialIndex, nextIndex int) (int, bool) {
	if nums[initialIndex] == nums[nextIndex] {
		nums[nextIndex] = 101
		return nextIndex, true
	} else {
		return nextIndex, false
	}
}

func sortArray(nums []int) []int {
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
