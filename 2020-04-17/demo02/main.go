package main

import "fmt"

func addTwo(nums []int, target int) []int {
	m := map[int]int{}

	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}

func main() {
	fmt.Println(addTwo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10))
	fmt.Println(addTwo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7))
}
