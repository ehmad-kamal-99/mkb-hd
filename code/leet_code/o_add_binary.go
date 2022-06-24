package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Problem Statement
// Given two binary strings a and b, return their sum as a binary string.

/*
Notes

The trick is to understand conversion of binary to decimal and then reconverting.
For ease, the func parseInt could be used for easy fix.
Try converting binary yourself.

Begin from the end.
Multiple that value with 2 raised to power counter (which is initialized to zero and gets incremented as the pointer moves)
Store the that value * 2 ^ counter into sum.
Repeat the process and keep adding to sum.
Return when last value encountered.

*/

func main() {
	// call function from here
	fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
		"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
	// fmt.Println(addBinary("1010", "1011"))
	// fmt.Println(addBinary("11", "1"))
	// fmt.Println(addBinary("0", "0"))
	// fmt.Println(addBinary("01", "00"))
}

func addBinary(a, b string) string {
	intA := convertToDecimal(a)
	intB := convertToDecimal(b)

	fmt.Println(intA)
	fmt.Println(intB)

	sum := intA + intB

	fmt.Println(sum)

	return strconv.FormatInt(int64(sum), 2)
}

func convertToDecimal(binary string) float64 {
	end := len(binary) - 1
	value := 0.0
	counter := 0.0

	for end > -1 {
		val, err := strconv.Atoi(string(binary[end]))
		if err != nil {
			panic(err)
		}

		value += float64(val) * math.Pow(2.0, counter)

		fmt.Println(value)

		counter++
		end--
	}

	return float64(value)
}

func convertToBinary(decimal int) string {
	binary := make([]string, 0)

	for decimal > 2 {
		binary = append(binary, strconv.Itoa(decimal%2))
		decimal = decimal / 2
	}

	binary = append(binary, strconv.Itoa(decimal))

	return strings.Join(binary, "")
}

// Challenge:
