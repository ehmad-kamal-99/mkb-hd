package main

import (
	"fmt"
)

// Problem Statement
// Given an integer array nums and an integer k, return the kth largest element in the array.
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
// You must solve it in O(n) time complexity.

/*
Notes

Use bubble-sort with outer loop running k-items and return k-1 item.

*/

func main() {
	// call function from here
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	//nums = sortKTimes(nums, k)
	nums = mergeSort(nums)

	fmt.Println(nums)

	return nums[k-1]
}

func sortKTimes(nums []int, k int) []int {
	for i := 0; i < k; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2

	return mergeArrays(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func mergeArrays(numsOne, numsTwo []int) []int {
	i, j, end := 0, 0, len(numsTwo)+len(numsOne)
	result := make([]int, end)

	for k := 0; k < len(result); k++ {
		if i == len(numsOne) && k < end {
			result[k] = numsTwo[j]
			j++
		} else if j == len(numsTwo) && k < end {
			result[k] = numsOne[i]
			i++
		} else if numsOne[i] > numsTwo[j] {
			result[k] = numsTwo[j]
			j++
		} else {
			result[k] = numsOne[i]
			i++
		}
	}

	return result
}

// Challenge: Solve it in O(N) time.
