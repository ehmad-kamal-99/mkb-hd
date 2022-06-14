package main

// Problem Statement
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

/*
Notes

1- Iterate over the array with 2 pointers
2- Keep adding values against the two pointers
3- if sum == target
	return indices of pointers

*/

func main() {
	// call function from here
	twoSum([]int{2, 7, 11, 15}, 9)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// Challenge: Less than O(n2) complexity.

func twoSumImproved(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for iter, num := range nums {
		hashmap[num] = iter
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if _, ok := hashmap[complement]; ok {
			if hashmap[complement] != i {
				return []int{hashmap[complement], i}
			}
		}
	}

	return []int{}
}
