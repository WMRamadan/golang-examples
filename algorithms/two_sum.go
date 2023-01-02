package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, ok := indexMap[complement]; ok {
			return []int{i, indexMap[complement]}
		}
		indexMap[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(nums, target))
}
