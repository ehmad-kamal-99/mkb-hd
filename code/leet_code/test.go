package main

import (
	"fmt"
)

func main() {
	// int sum(int N) {
	//    return N * (N + 1) / 2;
	// }

	//fmt.Println(27.0 / 26)
	//
	//test := []int{1, 2}
	//
	//fmt.Println(test[1:])
	//fmt.Println(test[1:2])
	//
	//fmt.Println(1 / 2)
	//
	//fmt.Println((math.Pow(10, 4) * (math.Pow(10, 4) + 1)) / 2)
	// 10 ^ 9 + 7

	var eightBit int8 = 56
	var eightBitSecond int8 = 56

	var result int16

	result = int16(eightBit) + 127

	fmt.Println(eightBit)
	fmt.Println(result)
	fmt.Println(eightBit + 127)
	fmt.Println(eightBit % 10)
	fmt.Println(eightBit % 107)
	fmt.Println((eightBit % 107) + (eightBitSecond % 107))
	fmt.Println((eightBit + eightBitSecond) % 107)
}
