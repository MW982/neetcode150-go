package main

import "fmt"

func maxProfit(prices []int) int {
	left := 0
	right := 1

	max_profit := 0
	for right < len(prices) {
		current_profit := prices[right] - prices[left]
		if current_profit > max_profit {
			max_profit = current_profit
		}
		if prices[right] < prices[left] {
			left = right
		}
		right++
	}

	return max_profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("hi", maxProfit(prices))

}
