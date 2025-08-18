package main

import "fmt"

// type Number interface {
// 	int | float64 | float32
// }
//
// type MinStack[T Number] struct {
// 	Stack []T
// 	Min   int
// }
//
// func Constructor[T Number]() MinStack[T] {
// 	return MinStack[T]{Min: 0, Stack: []T{}}
// }
//
// func (this *MinStack[T]) Push(val int) {
// }
//
// func (this *MinStack[T]) Pop() {
// }
//
// func (this *MinStack[T]) Top() T {
// 	return 1
// }
//
// func (this *MinStack[T]) GetMin() T {
// 	// return this.Min
// }
//

type MinStack struct {
	Stack []int
	Mins  []int
}

func Constructor() MinStack {
	return MinStack{Mins: []int{}, Stack: []int{}}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)

	if len(this.Mins) == 0 {
		this.Mins = append(this.Mins, val)
		return
	}

	this.Mins = append(this.Stack, min(this.GetMin(), val))
}

func (this *MinStack) Pop() {
	n := len(this.Stack) - 1
	this.Mins = this.Mins[:n]
	this.Stack = this.Stack[:n]
}

func (this *MinStack) Top() int {
	n := len(this.Stack)
	return this.Stack[n-1]
}

func (this *MinStack) GetMin() int {
	n := len(this.Mins)
	return this.Mins[n-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	fmt.Println("hi")
	obj := Constructor()
	obj.Push(7)
	fmt.Println(obj.GetMin())
	obj.Pop()

}
