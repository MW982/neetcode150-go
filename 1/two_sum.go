package main

import (
	"fmt"
)

func Abs(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func twoSum(nums []int, target int) []int {
	return []int{1}
}

func twoSumHashTableSolution(nums []int, target int) []int {
	// Returns two numbers rather than indexes
	var res []int
	hashTable := make(map[int]struct{})

	for _, num := range nums {
		hashTable[num] = struct{}{}
	}

	for _, num := range nums {
		val := Abs(target - num)
		fmt.Println('x', num, val, target)
		if _, hasNum := hashTable[val]; hasNum {
			return []int{num, val}
		}
	}

	fmt.Println("works", hashTable)
	return res
}

func twoSumHashTableIndexSolution(nums []int, target int) []int {
	hashTable := make(map[int]int)

	for index, num := range nums {
		diff := target - num
		if secondIndex, hasNum := hashTable[diff]; hasNum {
			return []int{index, secondIndex}
		}
		hashTable[num] = index
	}

	fmt.Println("works", hashTable)
	return []int{}
}

func main() {
	numbers := []int{1, 3, 5, 11, 9, 2}
	// fmt.Println("x", twoSumHashTableSolution(numbers, 5))
	fmt.Println("x", twoSumHashTableIndexSolution(numbers, 5))
	numbers = []int{1, 3}
	fmt.Println("x1", twoSumHashTableIndexSolution(numbers, 6))
	// fmt.Println("x", twoSum(numbers, 5))
	// fmt.Println("x", twoSumHashTableSolution(numbers, 123))
}
