package main

import (
	"fmt"
)

// Problem Statement
// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return
// the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.

/*
Notes

Basically divide and conquer, if that's what it means.
Start with the middle index, drop the irrelevant.

*/

func main() {
	// fmt.Println(1 / 2)
	// fmt.Println(2 / 2)
	// fmt.Println(3 / 2)
	// fmt.Println(4 / 2)

	// call function from here
	fmt.Println(searchInsertPosition([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsertPosition([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsertPosition([]int{1, 3, 5, 6}, 7))
}

func searchInsertPosition(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		}

		return 1
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}

		if i == len(nums)-1 {
			return len(nums)
		}
	}

	return -1
}

// func searchInsertPositionLogN(nums []int, target int) int {
// 	if len(nums) == 1 {
// 		if nums[0] >= target {
// 			return 0
// 		}
//
// 		return 1
// 	}
//
// 	left := 0
// 	right := len(nums) - 1
//
// 	for left <= right {
// 		mid := (left + right) / 2
//
// 		if nums[mid] == target {
// 			return mid
// 		} else if nums[mid] < target {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
//
// 	return left
// }

func searchInsertPositionLogN(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		}

		return 1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

// func sliceArray(nums []int) []int {
//
// }

// Challenge: LogN implementation
