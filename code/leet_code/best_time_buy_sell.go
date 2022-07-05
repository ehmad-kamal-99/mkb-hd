package main

import (
	"fmt"
)

// Problem Statement
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

/*
Notes

Brute Forcy

Pick first, compare with next, if bigger, update max profit
return max profit

[7, 10, 1, 3]

*/

func main() {
	// call function from here
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfitMine(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	if len(prices) == 2 {
		if prices[0] >= prices[1] {
			return 0
		}

		return prices[1] - prices[0]
	}

	maxProfit := 0

	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				if prices[j]-prices[i] > maxProfit {
					maxProfit = prices[j] - prices[i]
				}
			}
		}
	}

	return maxProfit
}

func maxProfit(prices []int) int {
	lowestBuyPrice := prices[0]
	profitIfSoldToday := 0
	maxProfit := 0

	for _, priceToday := range prices {
		if priceToday < lowestBuyPrice {
			lowestBuyPrice = priceToday
		}

		profitIfSoldToday = priceToday - lowestBuyPrice
		if maxProfit < profitIfSoldToday {
			maxProfit = profitIfSoldToday
		}
	}

	return maxProfit
}

// Challenge:
