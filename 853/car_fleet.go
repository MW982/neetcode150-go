package main

import (
	"fmt"
	"math"
	"sort"
)

type Car struct {
	Position int
	Velocity int
}

func floatGreaterOrEqual(a, b, eps float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return a > b
}

func (c1 Car) CanCatchUp(c2 Car, target int) bool {
	if c1.Velocity < c2.Velocity {
		return false
	}
	c1DistToTarget := target - c1.Position
	c2DistToTarget := target - c2.Position
	t1 := float64(c1DistToTarget) / float64(c1.Velocity)
	t2 := float64(c2DistToTarget) / float64(c2.Velocity)
	// fmt.Println(c1DistToTarget, c2DistToTarget, t1, t2)
	// if !floatGreaterOrEqual(t2, t1, 1e-9) {
	// 	return false
	// }

	return floatGreaterOrEqual(t2, t1, 1e-9)
}

func carFleet(target int, position []int, speed []int) int {
	// if len(position) == 1 {
	// 	return 1
	// }

	cars := []Car{}
	for index := range len(speed) {
		cars = append(cars, Car{Position: position[index],
			Velocity: speed[index]})
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Position > cars[j].Position
	})

	stack := []Car{}
	// numOfFleets := 0

	for _, car := range cars {
		if len(stack) == 0 {
			stack = append(stack, car)
			continue
		}

		topStack := stack[len(stack)-1]
		if car.CanCatchUp(topStack, target) {
			// topStack := stack[len(stack)-1]
			// stack = append(stack, car)
			continue
		} else {
			stack = append(stack, car)
			// stack = []Car{topStack}
			// numOfFleets++
		}

	}

	return len(stack)
}

func main() {
	position := []int{10, 8, 0, 5, 3}
	velocity := []int{2, 4, 1, 1, 3}
	target := 12

	fmt.Println("hi", carFleet(target, position, velocity))

	position = []int{3}
	velocity = []int{3}
	target = 10

	fmt.Println("hi", carFleet(target, position, velocity))

	position = []int{0, 4, 2}
	velocity = []int{2, 1, 3}
	target = 10

	fmt.Println("hi", carFleet(target, position, velocity))
}
