package main

import (
	"fmt"
	"strings"
)

func createHashKey(str string) string {
	hashMapKey := make(map[string]int)
	for _, char := range str {
		hashMapKey[string(char)]++
	}
	return convertHashKeyToString(hashMapKey)
}

func convertHashKeyToString(hash map[string]int) string {
	key := strings.Builder{}
	order := "abcdefghijklmnopqrstuvwxyz"

	for _, char := range order {
		value, hasChar := hash[string(char)]
		if hasChar {
			key.WriteString(fmt.Sprintf("%d%c", value, char))
		}
	}

	return key.String()
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	set := make(map[string][]string)
	for _, str := range strs {
		key := createHashKey(str)
		val, hasKey := set[key]
		if hasKey {
			set[key] = append(val, str)
		} else {
			set[key] = []string{str}
		}
	}
	fmt.Println(set)

	for _, value := range set {
		result = append(result, value)
	}
	return result
}

func main() {
	example := []string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}
	fmt.Println(groupAnagrams(example))
	// for _, word := range example {
	// 	sort.Slice(word, func(i, j int) bool {
	// 		return word[i] < word[j]
	// 	})
	// 	fmt.Println(string(word))
	// }

	// fmt.Println(string(example))
}
