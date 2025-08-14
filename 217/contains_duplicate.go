package main

import "fmt"

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, hasNum := set[num]; hasNum {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 3, 5, 10, 11, 1}
	fmt.Println("HEllo", containsDuplicate(nums))
	//        nums1 := [3]int{1, 3, }
	//        nums1[1] = 11111
	// fmt.Println("HElo", nums1 )
}
