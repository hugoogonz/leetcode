package main

import "fmt"

func twoSum(nums []int, target int) []int {

	newMap := make(map[int]int)

	for i, value := range nums {
		c := target - value
		if index, ok := newMap[c]; ok {
			return []int{index, i}
		}
		newMap[value] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{7, 6, 5, 3, 2, 1, 4, 9, 10}, 17))
}
