package main

import "fmt"

func twoSum(numbers []int, target int) []int {

	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum > target {
			right--
		} else {
			left++

		}
	}

	return []int{}
}

func main() {

	nums := []int{2, 7, 11, 15}
	fmt.Println("hi 1, 2", twoSum(nums, 9))
	fmt.Println("hi 2 3", twoSum(nums, 18))
	fmt.Println("hi 1 3", twoSum(nums, 13))
	fmt.Println("hi 3 4", twoSum(nums, 26))

}
