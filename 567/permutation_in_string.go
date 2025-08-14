package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Key := [26]int{}
	for _, char := range s1 {
		s1Key[char-'a']++
	}

	left, right := 0, len(s1)-1
	s2Key := [26]int{}
	// fmt.Println(left, right)
	for _, char := range s2[left : right+1] {
		s2Key[char-'a']++
	}

	for right < len(s2) {
		if s2Key == s1Key {
			return true
		} else {
			// fmt.Println(s2Key, s1Key, s2[left:right+1])
			leftChar := s2[left]
			s2Key[leftChar-'a']--
			left++
			right++
			if right >= len(s2) {
				break
			}
			rightChar := s2[right]
			s2Key[rightChar-'a']++
		}

	}

	return false
}

func main() {
	s1, s2 := "ab", "eidbaooo"
	fmt.Println("hi", checkInclusion(s1, s2))
	s1, s2 = "ab", "e"
	fmt.Println("hi", checkInclusion(s1, s2))
	s1, s2 = "a", "a"
	fmt.Println("hi", checkInclusion(s1, s2))
}
