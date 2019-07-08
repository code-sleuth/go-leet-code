package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		idx, ok := m[target-v]
		if ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
