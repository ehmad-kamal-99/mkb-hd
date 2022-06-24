package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0.0

	fmt.Println(math.Pow(10, 11))

	for i := 1.0; i < math.Pow(10, 11); i++ {
		sum += i
	}

	fmt.Println(sum)
}
