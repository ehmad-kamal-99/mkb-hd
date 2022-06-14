package main

import (
	"fmt"
)

// Problem Statement
// You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.
// Increment the large integer by one and return the resulting array of digits.

/*
Notes

I mean, pick the last one, update it & return the same array?

*/

func main() {
	// call function from here
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
	fmt.Println(plusOne([]int{1, 0, 0, 9}))
	fmt.Println(plusOne([]int{1, 0, 9, 9}))
	fmt.Println(plusOne([]int{1, 9, 9, 9}))
}

func plusOne(digits []int) []int {
	j := len(digits) - 1

	for {
		if digits[j] == 9 {
			digits[j] = 0

			j--

			if j == 0 {
				if digits[j] == 9 {
					newArr := []int{1}

					newArr = append(newArr, digits...)

					return newArr
				}

				digits[j] += 1

				break
			} else if j == -1 {
				newArr := []int{1}

				newArr = append(newArr, digits...)

				return newArr
			}

			break
		}

		digits[j] += 1

		break
	}

	return digits
}

// Challenge:
