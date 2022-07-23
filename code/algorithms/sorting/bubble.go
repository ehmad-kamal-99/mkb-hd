package main

import (
	"fmt"
)

/*
Normal Bubble Sort

[1, 2, 3, 4, 6, 5]

[1, 2, 3, 4, 6, 5]
 ^  ^

[1, 2, 3, 4, 6, 5]
 ^     ^

[1, 2, 3, 4, 6, 5]
 ^        ^

Improved Bubble Sort

[1, 2, 3, 4, 6, 5]

[1, 2, 3, 4, 6, 5]
 ^  ^

[1, 2, 3, 4, 6, 5]
    ^  ^

[1, 2, 3, 4, 6, 5]
       ^  ^

[1, 2, 3, 4, 6, 5]
          ^  ^

[1, 2, 3, 4, 6, 5]
             ^  ^

[1, 2, 3, 4, 6, 5]
 ^  ^

*/

func main() {
	test := []int{1, 2, 3, 0, 0, 0, 2, 5, 6}
	test2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort(test)
	fmt.Println(test)
	BubbleSortImproved(test)
	fmt.Println(test)
	BubbleSortImproved(test2)
	fmt.Println(test2)
}

func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
			}
		}
	}
}

func BubbleSortImproved(nums []int) {
	for {
		swap := false
		i := 0
		j := i + 1

		for j < len(nums) {
			if nums[i] > nums[j] {
				nums[j], nums[i] = nums[i], nums[j]
				swap = true
			}

			i++
			j++
		}

		if !swap {
			break
		}
	}
}
