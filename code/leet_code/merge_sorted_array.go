package main

import (
	"fmt"
)

// Problem Statement
// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

/*
Notes

I solved it. But now Imma repeat it for array :(

*/

func main() {
	// call function from here
	test1 := []int{1, 2, 3, 0, 0, 0}
	test2 := []int{2, 5, 6}

	merge(test1, 3, test2, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 0 {
		if len(nums2) == 0 {
			return
		}

		nums1 = append(nums1, nums2...)

		return
	}

	if len(nums2) == 0 {
		return
	}

	nums1 = append(nums1, nums2...)

	nums1 = bubbleSort(nums1)

	counter := 0

	for i := 0; i < len(nums1); i++ {
		if nums1[i] == 0 {
			counter++
		}
	}

	if counter > 0 {
		nums1 = nums1[counter:]
	}

	fmt.Println(nums1)
}

func bubbleSort(nums1 []int) []int {
	for i := 0; i < len(nums1)-1; i++ {
		for j := i + 1; j < len(nums1); j++ {
			if nums1[i] > nums1[j] {
				temp := nums1[j]
				nums1[j] = nums1[i]
				nums1[i] = temp
			}
		}
	}

	return nums1
}

// Challenge:
