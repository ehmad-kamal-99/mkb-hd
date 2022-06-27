package main

import (
	"fmt"
)

// Problem Statement
// Given an array, rotate the array to the right by k steps, where k is non-negative.

/*
Notes



*/

func main() {
	// call function from here
	test := []int{1, 2, 3, 4, 5, 6}
	test2 := []int{-1, -100, 3, 99}
	test3 := []int{1, 2}
	rotate(test, 11)
	rotate(test2, 2)
	rotate(test3, 3)
	fmt.Println(test)
	fmt.Println(test2)
	fmt.Println(test3)
}

func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}

	if k > len(nums) {
		if k%len(nums) == 0 {
			return
		}

		k = k % len(nums)
	}

	mapState := make(map[int]int)

	for i, val := range nums {
		mapState[i] = val
	}

	for i := 0; i < len(nums); i++ {
		nums[newIndex(i, k, len(nums))] = mapState[i]
	}
}

func newIndex(index, shift, len int) int {
	if index+shift >= len {
		return (index + shift) - len
	}

	return index + shift
}

// Challenge: Do it in-place with O(1) space complexity.
// Challenge: Do it in-place with O(n) time complexity.
