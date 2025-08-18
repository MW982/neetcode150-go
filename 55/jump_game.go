package main

import "fmt"

func canJump(nums []int) bool {
	goalIndex := len(nums) - 1
	currentIndex := goalIndex - 1
	for currentIndex >= 0 {
		if (goalIndex - currentIndex) <= nums[currentIndex] {
			goalIndex = currentIndex
		}
		currentIndex--
	}

	return goalIndex == 0
}

func canJumpOtherSolution(nums []int) bool {
	gas := 0
	for _, value := range nums {
		if gas < 0 {
			return false
		}
		if value > gas {
			gas = value
		}
		gas--
	}

	return true
}

func main() {
	fmt.Println("hi")
	nums := []int{2, 2, 2, 0, 4}
	fmt.Println(canJumpOtherSolution(nums))
}
