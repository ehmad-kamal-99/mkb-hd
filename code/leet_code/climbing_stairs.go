package main

import (
	"fmt"
)

// Problem Statement
// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

/*
Notes

One way to climb is to climb one step only. Then you can switch out each 2 single steps with one step.
Keep going on and that will lead to a lot of possible ways..
4 stairs
1-1-1-1
1-1-2
1-2-1
2-1-1
2-2

5 stairs
1-1-1-1-1
1-1-1-2
1-1-2-1
1-2-1-1
2-1-1-1
2-1-2
2-2-1
1-2-2

6 stairs
1-1-1-1-1-1
1-1-1-1-2
1-1-1-2-1
1-1-2-1-1
1-2-1-1-1
2-1-1-1-1
2-1-1-2
2-1-2-1
2-2-1-1
2-2-2
1-1-2-2
1-2-1-2
2-1-1-2
2-1-2-1
2-2-1-1

*/

func main() {
	// call function from here
	fmt.Println(climbingStairs(2))
	fmt.Println(climbingStairs(3))
	fmt.Println(climbingStairs(4))
	fmt.Println(climbingStairs(5))
	fmt.Println(climbingStairs(6))
}

func climbingStairs(n int) int {
	i, one, two := 0, 1, 1

	for i < n-1 {
		one, two = one+two, one
		i++
	}

	return one
}

// Challenge:
