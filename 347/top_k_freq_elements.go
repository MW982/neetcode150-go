package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	n := len(nums)

	for _, num := range nums {
		counter[num]++
	}

	buckets := make([][]int, n+1)
	for num, freq := range counter {
		buckets[freq] = append(buckets[freq], num)
	}

	result := []int{}
	fmt.Println(buckets)
	for i := n; i >= 0; i-- {
		for _, value := range buckets[i] {
			if k == 0 {
				return result
			}
			fmt.Println(value)
			result = append(result, value)
			k -= 1
		}
	}

	return result
}

func main() {
	fmt.Println("test")
	nums := []int{1, 2, 2, 3, 3, 3, 4, 4, 9, 9, 9, 10, 10, 10, 11, 11, 11, 11}
	k := 5

	fmt.Println(topKFrequent(nums, k))

}
