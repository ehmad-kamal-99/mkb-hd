package main

import "fmt"

func main() {
	// Test 1

	arr := [2]string{"Hello", "世界"} // Array declared and initiated.

	testArr := arr // copied array to a new variable.

	testArr[0] = "Konichiwa!" // new variable modifies array.

	fmt.Println("original array: ", arr) // original array unaffected.
	fmt.Println("modified array: ", testArr)

	slc := []string{"Hello", "世界"} // Slice declared and initiated.

	testSlc := slc // a new variable pointing to same slice

	testSlc[0] = "Konichiwa!" // new slice updated.

	fmt.Println("original slice: ", slc) // original slice affected.
	fmt.Println("modified slice: ", testSlc)

	// Test 2

	coolSlice := make([]string, 3, 5)
	coolSlice[0], coolSlice[1] = "Hello", "World!" // Slice declared and initiated.
	fmt.Println(coolSlice)                         // original slice.

	updateSlice(coolSlice)
	fmt.Println(coolSlice) // original slice affected.

	appendToSlice(coolSlice)
	fmt.Println(coolSlice) // original slice unaffected.

	coolSlice[2] = "Wow"
}

func updateSlice(slice []string) {
	slice[0] = "Bye"
}

func appendToSlice(slice []string) {
	slice = append(slice, "Foreva")
}
