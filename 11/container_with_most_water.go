package main

import "fmt"

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxAmountOfWater := 0
	for left < right {
		leftHeight := height[left]
		rightHeight := height[right]

		water := min(leftHeight, rightHeight) * (right - left)
		if water > maxAmountOfWater {
			maxAmountOfWater = water
		}
		if leftHeight <= rightHeight {
			left++
		} else {
			right--
		}
	}

	return maxAmountOfWater
}

func main() {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("hi", maxArea(heights))

}
