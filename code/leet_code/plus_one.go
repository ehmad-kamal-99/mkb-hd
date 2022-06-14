package main

import (
	"fmt"
)

// Problem Statement
// You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.
// Increment the large integer by one and return the resulting array of digits.

/*
Notes

I mean, pick the last one, update it & return the same array? No, what if the last int is 9, then you need to carry one and update the next int.
Add a check if array size is 1.
If array size equals 2, check the last int, if its 9, update it 0 and check the next int.
If that is 9 too, update that to zero until the array is empty.
Recursion? Let's avoid that.
How about we start looping over from last digit.
If that digit is not 9, add one and break?
If it is 9, make it 0 & keep moving until end of list.
Now if first item is 0, add one to beginning and return the array..

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
	if len(digits) == 1 {
		if digits[0] == 9 {
			return []int{1, 0}
		}

		digits[0] = digits[0] + 1

		return digits
	}

	end := len(digits) - 1

	for end > -1 {
		if digits[end] == 9 {
			digits[end] = 0
			end--

			continue
		}

		digits[end] = digits[end] + 1

		break
	}

	if digits[0] == 0 {
		newDigits := []int{1}

		newDigits = append(newDigits, digits...)

		return newDigits
	}

	return digits
}

// Challenge:
