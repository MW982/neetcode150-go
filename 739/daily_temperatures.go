package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	stack := [][2]int{}
	for index, value := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, [2]int{value, index})
			continue
		}

		n := len(stack)
		topOfStack, i := stack[n-1][0], stack[n-1][1]

		for topOfStack < value {
			// fmt.Println(stack, stack[len(stack)-1][0], value)
			daysDiff := index - i
			result[i] = daysDiff
			stack = stack[:n-1]
			n = len(stack)
			if n == 0 {
				break
			}
			topOfStack, i = stack[n-1][0], stack[n-1][1]
		}

		stack = append(stack, [2]int{value, index})

	}

	return result
}

func main() {
	fmt.Println("hi")
	temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println("hi", dailyTemperatures(temps))
	temps = []int{30, 60, 90}
	fmt.Println("hi", dailyTemperatures(temps))
	temps = []int{30, 40, 50, 60}
	fmt.Println("hi", dailyTemperatures(temps))
}
