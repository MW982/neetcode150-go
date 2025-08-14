package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	setTest := make(map[int]int)

	for _, value := range s {
		// fmt.Printf("%d %c\n", index, value)
		_, hasNum := setTest[int(value)]
		// fmt.Println(x,hasNum, int(value))
		if hasNum {
			setTest[int(value)] += 1
		} else {
			setTest[int(value)] = 1
		}
	}
	// fmt.Println("Rar", setTest)

	// for k,v := range setTest {
	// 	fmt.Println("x", k,v )
	// }

	for _, value := range t {
		_, hasNum := setTest[int(value)]
		// fmt.Println(x,hasNum, int(value))

		if hasNum {
			setTest[int(value)] -= 1
			if setTest[int(value)] == -1 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	ana := "anagram"
	gram := "nagaram"

	fmt.Println("HEllo", isAnagram(ana, gram))
}
