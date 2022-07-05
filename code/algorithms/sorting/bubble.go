package main

import (
	"fmt"
)

/*
Best case = O(N)

[1,2,3]

i = 1
j = 2

compare, already in place

i = 1
j = 3

compare, already in place

i = 2
j = 3

compare, already in place

finish

[1, 3, 2]

i = 1
j = 3

compare, already in place

i = 1
j = 2

compare, already in place

i = 3
j = 2

swap

finish
*/

func main() {
	test := []int{1, 2, 3, 0, 0, 0, 2, 5, 6}
	BubbleSort(test)
	fmt.Println(test)
}

func BubbleSort(nums []int) {
	// TODO: Fix me. In best case, it will still give O(n^2). Update the algorithm so if there are no swaps, exit.
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
