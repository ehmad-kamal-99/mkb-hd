package main

import (
	"fmt"
)

func main() {
	//arrOne := []int{1, 2, 3, 7, 9}
	//arrTwo := []int{2, 4, 5, 6, 8, 10, 11}
	//
	//sortedArr := mergeTwoSortedArrays(arrOne, arrTwo)
	//fmt.Println(sortedArr)
	//
	//sortedArrTwo := mergeTwoSortedArrays(arrTwo, arrOne)
	//fmt.Println(sortedArrTwo)

	//arrThree := []int{2, 4, 5, 7, 9, 6, 8, 10, 11}
	//
	//mid := (len(arrThree)) / 2
	//
	//sortedArrThree := mergeOneSortedArrays(arrThree[:mid], arrThree[mid:])
	//fmt.Println(sortedArrThree)

	//arrFour := []int{2, 4, 5, 7, 8, 9, 1, 3, 6, 10}
	//
	//mid := (len(arrFour)) / 2
	//
	//sortedArrFour := mergeBasedOnIndex(arrFour, 0, mid, len(arrFour))
	//fmt.Println(sortedArrFour)
	//
	//arrFive := []int{2, 4, 5, 7, 8, 9, 1, 3, 6}
	//
	//mid := (len(arrFive)) / 2
	//
	//sortedArrFive := mergeBasedOnIndex(arrFour, 0, mid, len(arrFive))
	//fmt.Println(sortedArrFive)

	arr := []int{9, 8, 1, 3, 10, 24, 6, 4, 0}

	sortedArr := mergeSort(arr)

	fmt.Println(sortedArr)
}

func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	mid := (len(slice)) / 2

	return mergeTwoSortedArraysImproved(mergeSort(slice[:mid]), mergeSort(slice[mid:]))
}

//func mergeSort(arr []int, l, h int) []int {
//	if l < h {
//		mid := (l + h) / 2
//
//		mergeSort(arr, l, mid)
//		mergeSort(arr, mid, h)
//		arr = mergeBasedOnIndex(arr, l, mid, h)
//	}
//
//	return arr
//}

//
//func merge(l, m, h int) {
//
//}

//func mergeInPlace(arr []int, low, mid, high int) {
//	i, j := low, mid
//
//	for i < mid && j < high {
//		if arr[i] > arr[j] {
//			arr[i], arr[j] = arr[j], arr[i]
//			arr = append(arr, arr[j])
//			i++
//		} else if arr[i] < arr[j] {
//			arr = append(arr, arr[i])
//			i++
//		} else {
//			arr = append(arr, arr[i])
//			arr = append(arr, arr[j])
//			i++
//			j++
//		}
//	}
//
//	if i < mid {
//		arr = append(arr, arr[i:]...)
//	}
//
//	if j < high {
//		arr = append(arr, arr[j:]...)
//	}
//
//	return
//}

func mergeBasedOnIndex(arr []int, low, mid, high int) (mergedArr []int) {
	i, j := low, mid

	for i < mid && j < high {
		if arr[i] > arr[j] {
			mergedArr = append(mergedArr, arr[j])
			j++
		} else if arr[i] < arr[j] {
			mergedArr = append(mergedArr, arr[i])
			i++
		} else {
			mergedArr = append(mergedArr, arr[i])
			mergedArr = append(mergedArr, arr[j])
			i++
			j++
		}
	}

	if i < mid {
		mergedArr = append(mergedArr, arr[i:]...)
	}

	if j < high {
		mergedArr = append(mergedArr, arr[j:]...)
	}

	return
}

func mergeOneSortedArrays(arrOne, arrTwo []int) (mergedArr []int) {
	i, j := 0, 0

	for i < len(arrOne) && j < len(arrTwo) {
		if arrOne[i] > arrTwo[j] {
			mergedArr = append(mergedArr, arrTwo[j])
			j++
		} else if arrOne[i] < arrTwo[j] {
			mergedArr = append(mergedArr, arrOne[i])
			i++
		} else {
			mergedArr = append(mergedArr, arrOne[i])
			mergedArr = append(mergedArr, arrTwo[j])
			i++
			j++
		}
	}

	if i < len(arrOne) {
		mergedArr = append(mergedArr, arrOne[i:]...)
	}

	if j < len(arrTwo) {
		mergedArr = append(mergedArr, arrTwo[j:]...)
	}

	return
}

func mergeTwoSortedArrays(arrOne, arrTwo []int) (mergedArr []int) {
	i, j := 0, 0

	for i < len(arrOne) && j < len(arrTwo) {
		if arrOne[i] > arrTwo[j] {
			mergedArr = append(mergedArr, arrTwo[j])
			j++
		} else if arrOne[i] < arrTwo[j] {
			mergedArr = append(mergedArr, arrOne[i])
			i++
		} else {
			mergedArr = append(mergedArr, arrOne[i])
			mergedArr = append(mergedArr, arrTwo[j])
			i++
			j++
		}
	}

	if i <= len(arrOne) {
		mergedArr = append(mergedArr, arrOne[i:]...)
	}

	if j <= len(arrTwo) {
		mergedArr = append(mergedArr, arrTwo[j:]...)
	}

	return
}

func mergeTwoSortedArraysImproved(arrOne, arrTwo []int) []int {
	mergedArr, i, j := make([]int, len(arrOne)+len(arrTwo)), 0, 0

	for k := 0; k < len(mergedArr); k++ {
		if i == len(arrOne) && j < len(arrTwo) {
			mergedArr[k] = arrTwo[j]
			j++
		} else if i < len(arrOne) && j == len(arrTwo) {
			mergedArr[k] = arrOne[i]
			i++
		} else if arrOne[i] > arrTwo[j] {
			mergedArr[k] = arrTwo[j]
			j++
		} else {
			mergedArr[k] = arrOne[i]
			i++
		}
	}

	return mergedArr
}
