package main

import (
	"fmt"
	"strconv"
)

func isOperand(token string) bool {
	return isDiv(token) || isMul(token) || isAdd(token) || isSub(token)
}

func isDiv(token string) bool {
	return token == "/"
}

func isAdd(token string) bool {
	return token == "+"
}

func isMul(token string) bool {
	return token == "*"
}

func isSub(token string) bool {
	return token == "-"
}

func evalRPN(tokens []string) int {
	// size := len(tokens)
	stack := []string{}
	for _, token := range tokens {
		if isOperand(token) {
			n := len(stack) - 1
			left, _ := strconv.Atoi(stack[n-1])
			right, _ := strconv.Atoi(stack[n])

			stack = stack[:n-1]

			var res int
			if isSub(token) {
				res = left - right
			}
			if isDiv(token) {
				res = left / right
			}
			if isAdd(token) {
				res = left + right
			}
			if isMul(token) {
				res = left * right
			}
			stack = append(stack, strconv.Itoa(res))
		} else {
			stack = append(stack, token)
		}
	}
	result, _ := strconv.Atoi(stack[0])
	return result
}

func main() {
	fmt.Println("hi")
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))

	tokens = []string{"4", "13", "5", "/", "+"}

	fmt.Println(evalRPN(tokens))
	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))

}
