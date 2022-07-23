package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 76}
	fmt.Println(binarySearch(arr, 1))
	fmt.Println(binarySearch(arr, 2))
	fmt.Println(binarySearch(arr, 3))
	fmt.Println(binarySearch(arr, 4))
	fmt.Println(binarySearch(arr, 5))
	fmt.Println(binarySearch(arr, 6))
	fmt.Println(binarySearch(arr, 7))
	fmt.Println(binarySearch(arr, 8))
	fmt.Println(binarySearch(arr, 76))
	fmt.Println(binarySearch(arr, 100))
	fmt.Println(binarySearch(arr, 1000))
	fmt.Println(binarySearch(arr, 0))
}

func binarySearch(arr []int, item int) int {
	low, high := 0, len(arr)

	for {
		mid := (low + high) / 2

		if item > arr[mid] {
			low = mid

			if low == len(arr)-1 {
				return -1
			}

		} else if item < arr[mid] {
			high = mid

			if high == 0 {
				return -1
			}
		} else {
			if item == arr[mid] {
				return mid
			}

			return -1
		}
	}
}
