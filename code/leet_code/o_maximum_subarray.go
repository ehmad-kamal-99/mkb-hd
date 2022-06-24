package main

import (
	"fmt"
)

// Problem Statement
// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// A subarray is a contiguous part of an array.

/*
Notes

Brute Force way with O(n2) // failed *facepalm*

Start finding all permutations I believe. All the possible outcomes. Then add the values in the set. The highest value from a set wins.
If I have [5, 3, -4], I need to make sets like (5), (5,3), (5,3,-4), (3), (3,-4), (-4)

Lets take two pointers. Starting point of both is same. We increment one to keep increasing the size of array till the last index.
Then, we move the second pointer & repeat the process, until the first pointer reaches the end.
Each sublist needs to be persisted into another array/slice. Then, sum up all the slices/arrays and keep track of the biggest value.
That's the answer.

Runtime Error, needs fixing.
*/

func main() {
	// call function from here
	fmt.Println(maximumSubarray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maximumSubarray([]int{1}))
	fmt.Println(maximumSubarray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maximumSubarray(nums []int) int {
	allSubArrays := make([][]int, 0)
	startOne, startTwo := 0, 1

	for startOne < len(nums) {
		for startTwo <= len(nums) {
			allSubArrays = append(allSubArrays, nums[startOne:startTwo])
			startTwo++
		}
		startOne++
		startTwo = startOne + 1
	}

	biggestSum := -10000

	for _, array := range allSubArrays {
		sum := 0
		for _, val := range array {
			sum += val
		}

		if sum > biggestSum {
			biggestSum = sum
		}
	}

	return biggestSum
}

// I dont understand what's happening here.
func maxSubArray(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	var minSum, sum int // sum , minSum == 0

	res := -1 * 1 << 31 // largest negative number

	for i := 0; i < len(nums); i++ {
		fmt.Println("\n\niteration: ", i+1)
		fmt.Println("value: ", nums[i])
		fmt.Println("sum: ", sum)
		fmt.Println("minSum: ", minSum)
		fmt.Println("res: ", res, "\n\n")

		sum += nums[i] // sum = sum + nums[i]

		fmt.Println("sum updated, sum: ", sum)

		if res < sum-minSum {
			fmt.Println("res < sum - minSum, res: ", sum-minSum)

			res = sum - minSum // if res < sum - minSum, res = sum - minSum
		}

		if minSum > sum {
			fmt.Println("minSum > sum, minSum: ", sum)

			minSum = sum // if minSum > sum, minSum = sum
		}
	}

	return res
}

func maximumSubarrayON(nums []int) int {
	return 0
}

// Challenge: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
